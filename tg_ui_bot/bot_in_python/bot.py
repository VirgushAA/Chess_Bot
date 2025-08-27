import io

from telegram import Update, InlineKeyboardButton, InlineKeyboardMarkup
from telegram.ext import ApplicationBuilder, CommandHandler, ContextTypes, CallbackQueryHandler
import sqlite3
import requests
import re
from PIL import Image, ImageDraw

BASE_URL = "http://localhost:8080"


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
    r = requests.post(f"{BASE_URL}/newgame")
    data = r.json()
    game_id = data["gameId"]
    active_sessions[game_id] = {"player_white": update.effective_user.id, "player_black": None}
    await update.message.reply_text(f"♟ Новая игра создана!\nGame ID: {game_id}")
    await send_board_image(update, context, data["state"]["Board"]["Board"])


async def chess_join_by_id(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    if not context.args:
        await update.message.reply_text('Пожалуйста введи ID игры, к которой хочешь присоединиться.')
        return
    if not (context.args[0] in active_sessions and active_sessions[context.args[0]]['player_black'] is None):
        await update.message.reply_text('Игры с таким ID не существует.')
    else:
        active_sessions[context.args[0]]['player_black'] = update.effective_user.id
        await context.bot.sendMessage(chat_id=active_sessions[context.args[0]]['player_white'],
                                      text='Ваш соперник присоединился к игре!')
        await update.message.reply_text(f"Теперь ты участвуешь в игре с ID {context.args[0]}!")
        await send_board_image(update, context)


async def chess_make_move(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    if not context.args:
        await update.message.reply_text('Пожалуйста введи ход в формате "e2e4"')
        return

    move = normalize_move(context.args[0])

    r = requests.post(f"{BASE_URL}/move")
    data = r.json()
    # gamestate = data["state"]
    await send_board_image(update, context, data["state"]["Board"]["Board"])


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


def render_pieces_to_board(img, board):
    square_size = img.width // 8  # размер одной клетки
    half_square = square_size // 2

    for i in range(len(board)):
        piece_value = board[i]
        if piece_value != 0:
            piece_type, piece_color = get_piece_type_color(piece_value)
            row = i / 8
            col = i % 8
            piece_img = Image.open(get_piece_filename(piece_type, piece_color)).resize((50, 50)).convert("RGBA")
            # x = int(col * (img.width // 8) + (img.width // 16))
            # y = int(row * (img.height // 8) + (img.height // 16))
            x = int(col * square_size + half_square - piece_img.width // 2)
            y = int(row * square_size + half_square - piece_img.height // 2)
            img.paste(piece_img, (x, y), mask=piece_img)
    return img


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


async def send_board_image(update: Update, context: ContextTypes.DEFAULT_TYPE, board=None) -> None:
    img = render_board()
    if board:
        img = render_pieces_to_board(img, board)

    with io.BytesIO() as output:
        img.save(output, 'PNG')
        output.seek(0)
        await update.message.reply_photo(photo=output)


def normalize_move(text: str) -> str:
    text = text.lower().strip()
    text = re.sub(r"[-_\s]", "", text)  # Remove separators
    text = re.sub(r"[^a-h1-8]", "", text)   # Only keep letters/numbers
    return text[:4]


def get_piece_type_color(piece) -> tuple[int, int]:
    # Возвращает тип и цвет фигуры в том порядке
    return piece & 0x7, (piece >> 4) & 0x1


if __name__ == "__main__":
    active_sessions = {}
    create_db()

    app = ApplicationBuilder().token("7806801443:AAHrpGLx1Gd1WJG6mHuHcB_wAr_cQbTTU6w").build()

    app.add_handler(CommandHandler('start', start))

    app.add_handler(CommandHandler('users', leaderboard))

    app.add_handler(CommandHandler("newgame", chess_new_game))

    app.add_handler(CommandHandler("move", chess_make_move))

    app.add_handler(CommandHandler("join", chess_join_by_id))

    app.run_polling()
