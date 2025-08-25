from telegram import Update, InlineKeyboardButton, InlineKeyboardMarkup
from telegram.ext import ApplicationBuilder, CommandHandler, ContextTypes, CallbackQueryHandler
import sqlite3, requests, re

BASE_URL = "http://localhost:8080"

def create_db():
    con = sqlite3.connect('users.db')
    cur = con.cursor()
    cur.execute(''' CREATE TABLE IF NOT EXISTS users ( user_id INTEGER PRIMARY KEY,
                                                           username TEXT,
                                                           first_name TEXT,
                                                           score INTEGER DEFAULT 0 ) ''')

    # cur.execute(''' CREATE TABLE IF NOT EXISTS games ( game_id INTEGER PRIMARY KEY,
    #                                                            username TEXT,
    #                                                            first_name TEXT,
    #                                                            score INTEGER DEFAULT 0 ) ''')

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
    r = requests.get(f"{BASE_URL}/newgame")
    data = r.json()
    game_id = data["gameId"]

    await update.message.reply_text(f"♟ Новая игра создана!\nGame ID: {game_id}")



async def chess_make_move(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    if not context.args:
        await update.message.reply_text('Пожалуйста введи ход в формате "e2e4"')
        return

    move = normalize_move(context.args[0])



def normalize_move(text: str) -> str:
    text = text.lower().strip()
    text = re.sub(r"[-_\s]", "", text)  # Remove separators
    text = re.sub(r"[^a-h1-8]", "", text)   # Only keep letters/numbers
    return text[:4]

if __name__ == "__main__":
    create_db()

    app = ApplicationBuilder().token("7806801443:AAHrpGLx1Gd1WJG6mHuHcB_wAr_cQbTTU6w").build()

    app.add_handler(CommandHandler('start', start))

    app.add_handler(CommandHandler('users', leaderboard))

    app.add_handler(CommandHandler("/newgame", chess_new_game))

    app.add_handler(CommandHandler("/move", chess_make_move))

    app.run_polling()
