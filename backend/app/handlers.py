from aiogram import F, Router
from aiogram.filters import CommandStart
from aiogram.types import Message, CallbackQuery

from app.keyboards import button_appeal, button_with_tems, choos_topic_in_appeal

router = Router()

appeals = {"1": "Ремонт телика", "2": "Ремонт компа", "3": "Ремонт стиралки", "4": "Ремонт утюга"}


# Регистрация и аунтефикация при старте бота(Админ/Пользователь)
@router.message(CommandStart())
async def start_handler(message: Message):
    if True:
        await message.answer(f'Вы авторизированы',
                             reply_markup=await button_appeal())


@router.message(F.text == 'Мой профиль ⚙️')
async def creat_list(message: Message):
    await message.answer(f'Выберите 👇', reply_markup=await button_with_tems())


@router.callback_query(F.data == 'Создать обращение')
async def creat_appeal(callback: CallbackQuery):
    await callback.answer()
    await callback.message.edit_text('Создать обращение 📝', reply_markup=await choos_topic_in_appeal())


@router.callback_query(F.data == 'Жалобы')
async def creat_appeal_complaint(callback: CallbackQuery):
    await callback.answer()
