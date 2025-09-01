import io
from tarfile import tar_filter

from telegram import Update, InlineKeyboardButton, InlineKeyboardMarkup
from telegram.ext import ApplicationBuilder, CommandHandler, ContextTypes, CallbackQueryHandler
import sqlite3
import requests
import re
from PIL import Image, ImageDraw, ImageFont

from dotenv import load_dotenv
import os

load_dotenv()

BOT_TOKEN = os.getenv('BOT_TOKEN')
BASE_URL = os.getenv('CHESS_API_URL')


def create_db():
    con = sqlite3.connect('users.db')
    cur = con.cursor()
    cur.execute(''' CREATE TABLE IF NOT EXISTS users ( user_id INTEGER PRIMARY KEY,
                                                           username TEXT,
                                                           first_name TEXT,
                                                           score INTEGER DEFAULT 0 ) ''')

    con.commit()
    con.close()


async def register_user(user: Update.effective_user):
    con = sqlite3.connect('users.db')
    cur = con.cursor()
    user_id = user.id
    user_name = user.username or ''
    first_name = user.first_name or ''
    # cur.execute('INSERT OR IGNORE INTO users VALUES (?, ?)', (user_id, 0))
    cur.execute('''
        INSERT INTO users (user_id, username, first_name)
        VALUES (?, ?, ?)
        ON CONFLICT(user_id) DO UPDATE SET
            username = excluded.username,
            first_name = excluded.first_name
    ''', (user_id, user_name, first_name))
    con.commit()
    con.close()


async def leaderboard(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    con = sqlite3.connect('users.db')
    cur = con.cursor()
    cur.execute('SELECT first_name, username, score FROM users ORDER BY score DESC')
    rows = cur.fetchall()
    con.close()

    if not rows:
        await update.message.reply_text("Пока нет участников. Ты будешь первым!")
        return

    leaderboard_text = "🏆 Таблица лидеров:\n\n"
    for idx, (first_name, username, score) in enumerate(rows, start=1):
        name = f"{first_name} (@{username})" if username else first_name
        leaderboard_text += f"{idx}. {name}: {score} побед\n"

    await update.message.reply_text(leaderboard_text)


async def start(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    await register_user(update.effective_user)
    await update.message.reply_photo(photo=open('images/xdd.jpeg', 'rb'), caption='Приветствую тебя в самом лучшем боте!')


async def chess_new_game(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    if find_game_with_user(update.effective_user.id):
        await update.message.reply_text("Нельзя создавать игру уже находясь в игре.")
        return

    r = requests.post(f"{BASE_URL}/newgame")
    data = r.json()
    game_id = data["gameId"]
    active_sessions[game_id] = {"player_white": update.effective_user.id, "player_black": None, "Turn": 0}
    await update.message.reply_text(f"♟ Новая игра создана!\nGame ID: {game_id}")
    await send_board_image(update, context, data["state"]['Board']['Board'],
                           (find_players_color_in_game(update.effective_user.id) + 1) % 2)
    print(active_sessions)


async def chess_join_by_id(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    if not context.args:
        await update.message.reply_text('Пожалуйста введи ID игры, к которой хочешь присоединиться.')
        return
    if not (context.args[0] in active_sessions):
        await update.message.reply_text('Игры с таким ID не существует.')
        return
    game_id = active_sessions[context.args[0]]
    user_game = find_game_with_user(update.effective_user.id)
    if user_game is not None and user_game != context.args[0]:
        await update.message.reply_text("Нельзя присоединиться к чужой игре, если ты уже в игре.")
        return

    if game_id['player_white'] is None:
        game_id['player_white'] = update.effective_user.id
        await context.bot.sendMessage(chat_id=active_sessions[context.args[0]]['player_white'],
                                      text='Ваш соперник присоединился к игре!')
        await update.message.reply_text(f"Теперь ты участвуешь в игре с ID {context.args[0]}!")
        await send_default_board_image(update, context, (find_players_color_in_game(update.effective_user.id) + 1) % 2
                                       if find_players_color_in_game(update.effective_user.id) != 'both' else 1)
        print(active_sessions)
    elif game_id['player_black'] is None:
        game_id['player_black'] = update.effective_user.id
        await context.bot.sendMessage(chat_id=active_sessions[context.args[0]]['player_white'],
                                      text='Ваш соперник присоединился к игре!')
        await update.message.reply_text(f"Теперь ты участвуешь в игре с ID {context.args[0]}!")
        await send_default_board_image(update, context, (find_players_color_in_game(update.effective_user.id) + 1) % 2
                                       if find_players_color_in_game(update.effective_user.id) != 'both' else 0)
        print(active_sessions)
    else:
        await update.message.reply_text(f"В это игре все места уже заняты!")


async def chess_game_over(update: Update, context: ContextTypes.DEFAULT_TYPE, state) ->None:
    game = find_game_with_user(update.effective_user.id)
    if not game:
        return

    if state["Stalemate"]:
        winner_text = "Ничья! 🤝"
    else:
        winner_color = (active_sessions[game]['Turn'] + 1) % 2
        winner_text = "Белые победили! 🏆" if winner_color == 0 else "Чёрные победили! 🏆"

        for player_id in [game.get('player_white'), game.get('player_black')]:
            if player_id:
                await context.bot.sendMessage(chat_id=player_id, text=winner_text)
                await chess_leave_game(update, context, player_id)


async def chess_leave_game(update: Update, context: ContextTypes.DEFAULT_TYPE, player=None) -> None:
    if player is None:
        player = update.effective_user.id
    game_id = find_game_with_user(player)
    game = active_sessions[game_id]
    requests.post(f"{BASE_URL}/endgame", json={'gameId': game_id})
    if game['player_white'] == player:
        game['player_white'] = None
    if game['player_black'] == player:
        game['player_black'] = None
    await context.bot.sendMessage(chat_id=player, text='Ты вышел из игры.')

    if game['player_white'] is None and game['player_black'] is None:
        _ = active_sessions.pop(game_id, None)

    print(active_sessions)


async def chess_make_move(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    if not context.args:
        await update.message.reply_text('Пожалуйста введи ход в формате "e2e4"')
        return
    game = find_game_with_user(update.effective_user.id)
    if isinstance(game, tuple):
        game = game[0]
    if not game:
        await update.message.reply_text('Но ты не в игре.')
        return
    if (active_sessions[game]['Turn'] != find_players_color_in_game(update.effective_user.id)
            and find_players_color_in_game(update.effective_user.id) != 'both'):
        await update.message.reply_text('Сейчас не твой ход.')
        return

    turn_in = {
        'gameId': game,
        'move': ''
    }

    move = normalize_move(context.args[0])
    print(move)
    turn_in['move'] = move
    r = requests.post(f"{BASE_URL}/move", json=turn_in)

    data = r.json()
    active_sessions[game]['Turn'] = (active_sessions[game]['Turn'] + 1) % 2
    await send_board_image(update, context, data["state"]['Board']['Board'], (data['state']['Turn'] + 1) % 2)

    if find_players_color_in_game(update.effective_user.id) != 'both':
        opponent_id = (active_sessions[game]['player_white']
                       if active_sessions[game]['player_black'] == update.effective_user.id
                       else active_sessions[game]['player_black'])
        if opponent_id is not None:
            await send_board_image(opponent_id, context, data["state"]['Board']['Board'], (data['state']['Turn'] + 1) % 2)

    print(active_sessions)


def render_board(width=400, height=400):
    square_size = width // 8

    img = Image.new('RGBA', (width, height), color='#FFFFFF')
    draw = ImageDraw.Draw(img)

    colors = ["#FFCE9E", "#D18B47"]

    for row in range(8):
        for col in range(8):
            color_idx = (row + col) % 2
            color = colors[color_idx]

            top_left = (col * square_size, row * square_size)
            bot_right = ((col + 1) * square_size, (row + 1) * square_size)

            draw.rectangle([top_left, bot_right], fill=color)

    return img


def render_pieces_to_board(img, board, color=0):
    square_size = img.width // 8

    for row in range(7, -1, -1):
        for col in range(8):
            i = row * 8 + col
            piece_value = board[i]

            if piece_value != 0:
                piece_type, piece_color = get_piece_type_color(piece_value)
                piece_img = Image.open(get_piece_filename(piece_type, piece_color)).resize((50, 50)).convert("RGBA")

                if color:
                    x = col * square_size
                    y = (7 - row) * square_size
                else:
                    x = (7 - col) * square_size
                    y = row * square_size
                img.paste(piece_img, (x, y), mask=piece_img)

    return img


# def render_labels_to_board(img, color=0):
#     draw = ImageDraw.Draw(img)
#     font = ImageFont.truetype("/System/Library/Fonts/AppleSDGothicNeo.ttc", 20)
#     square_size = img.width // 8

#             letter = chr(ord('A') + col)
#             draw.text((x, y), letter, font=font, fill='black', anchor='ms')
#
#     return img


def get_piece_filename(piece_type, piece_color):
    piece_types = {
        0: "empty",
        1: "P",
        2: "N",
        3: "B",
        4: "R",
        5: "Q",
        6: "K"
    }

    pieces_folder = 'images/pieces/'
    color_prefix = 'w' if piece_color == 0 else 'b'

    filename = pieces_folder + color_prefix + piece_types[piece_type] + '.png'
    return filename


async def send_board_image(target, context: ContextTypes.DEFAULT_TYPE, board=None, color=0) -> None:
    img = render_board()
    if board:
        img = render_pieces_to_board(img, board, color)

    with io.BytesIO() as output:
        img.save(output, 'PNG')
        output.seek(0)

        if isinstance(target, Update):
            await target.message.reply_photo(photo=output)
        else:
            await context.bot.sendPhoto(chat_id=target, photo=output)


async def send_default_board_image(target, context: ContextTypes.DEFAULT_TYPE, color=0) -> None:
    board = [4, 2, 3, 5, 6, 3, 2, 4, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
             0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 17, 17, 17, 17, 17, 17, 17, 17, 20, 18, 19, 21, 22, 19, 18, 20]

    img = render_board()
    img = render_pieces_to_board(img, board, color)

    with io.BytesIO() as output:
        img.save(output, 'PNG')
        output.seek(0)

        if isinstance(target, Update):
            await target.message.reply_photo(photo=output)
        else:
            await context.bot.sendPhoto(chat_id=target, photo=output)


def normalize_move(text: str) -> str:
    text = text.lower().strip()
    text = re.sub(r"[-_\s]", "", text)  # Remove separators
    text = re.sub(r"[^a-h1-8]", "", text)   # Only keep letters/numbers
    return text[:4]


def get_piece_type_color(piece) -> tuple[int, int]:
    # Возвращает тип и цвет фигуры в том порядке
    return piece & 0x7, (piece >> 4) & 0x1


def find_game_with_user(user_id):
    for game_id, players in active_sessions.items():
        if players.get('player_white') == user_id or players.get('player_black') == user_id:
            return game_id
    return None


def find_players_color_in_game(user_id):
    game = find_game_with_user(user_id)
    if game:
        if active_sessions[game]['player_white'] == user_id and active_sessions[game]['player_black'] == user_id:
            return 'both'
        elif active_sessions[game]['player_white'] == user_id:
            return 0
        elif active_sessions[game]['player_black'] == user_id:
            return 1

    return None


if __name__ == "__main__":
    active_sessions = {}
    create_db()

    app = ApplicationBuilder().token(BOT_TOKEN).build()

    app.add_handler(CommandHandler('start', start))

    app.add_handler(CommandHandler('users', leaderboard))

    app.add_handler(CommandHandler("newgame", chess_new_game))

    app.add_handler(CommandHandler("move", chess_make_move))

    app.add_handler(CommandHandler("join", chess_join_by_id))

    app.add_handler(CommandHandler("leave", chess_leave_game))

    app.run_polling()
