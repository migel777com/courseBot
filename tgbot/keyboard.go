package tgbot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

var courseStart = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Начать курс","/startCourse"),
	),
)

var NextModule = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Перейти к следующему модулю","/nextModule"),
	),
)

var mainMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Главное меню"),
	),
)

var moduleKeys = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Модуль 1. Эффективности можно научиться","/openModule 1"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Модуль 2. Сущность работы управляющих","/openModule 2"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Модуль 3. Можно ли научиться эффективности?","/openModule 3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Модуль 4. Знайте свое время","/openModule 4"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Модуль 5. Что означает регистрация времени?","/openModule 5"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Модуль 6. Что означает управление временем?","/openModule 6"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Модуль 7. Что означает укрупнение времени?","/openModule 7"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Модуль 8. Собственные обязательства управляющего","/openModule 8"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Модуль 9. Повышение эффективности работы специалиста ","/openModule 9"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Модуль 10. Правильные взаимоотношения","/openModule 10"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Модуль 11. Ставка на сильные качества","/openModule 11"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Модуль 12. Повышение собственной эффективности ","/openModule 12"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Модуль 13. Всему свое время ","/openModule 13"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Модуль 14. Приоритеты","/openModule 14"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Модуль 15. Элементы принятия решении","/openModule 15"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Модуль 16. Эффективные решения","/openModule 16"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Модуль 17. Заключение","/openModule 17"),
	),

)
