import yaml
import asyncio
from app.handlers import router
from aiogram import Bot, Dispatcher

with open("config.yaml", "r") as file:
    data = yaml.safe_load(file)

bot = Bot(token=data['TOKEN'])
dp = Dispatcher()


async def main():
    dp.include_router(router=router)
    await dp.start_polling(bot)


if __name__ == '__main__':
    asyncio.run(main())
