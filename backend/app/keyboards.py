from aiogram.types import InlineKeyboardMarkup, InlineKeyboardButton, ReplyKeyboardMarkup, KeyboardButton
from aiogram.utils.keyboard import InlineKeyboardBuilder


async def button_appeal():
    my_profil = ReplyKeyboardMarkup(keyboard=[
        [KeyboardButton(text='Мой профиль ⚙️')]
    ], resize_keyboard=True)
    return my_profil


async def button_with_tems():
    array_tems = ['Создать обращение', 'Посмотреть обращения', 'Посмотреть статус обращения']
    tems = InlineKeyboardBuilder()
    for tem in array_tems:
        tems.add(InlineKeyboardButton(text=tem, callback_data=tem))
    return tems.adjust(1).as_markup()


async def choos_topic_in_appeal():
    array_topic = ['Жалобы', 'Запросы на ремонт', 'Предложения']
    topic = InlineKeyboardBuilder()
    for top in array_topic:
        topic.add(InlineKeyboardButton(text=top, callback_data=top))
    return topic.adjust(2).as_markup()



