package tgbot

import (
	"courseBot/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sethvargo/go-password/password"
	"log"
	"strconv"
	"strings"
)

func HandleCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI, db *database.DBAccess) {
	server := Server{Bot: bot, Db: db}

	chatID := update.Message.Chat.ID
	message := update.Message.Text

	if !server.Db.CheckUser(chatID) {
		if server.Db.CheckPassword(update) {
			msg := tgbotapi.NewMessage(chatID, "Что бы перейти на следующий модуль, сделайте все домашние задания")
			msg.ReplyMarkup = courseStart
			_, _ = server.Bot.Send(msg)
			return
		} else if message == "/start" {
			msg := tgbotapi.NewMessage(chatID, "Введите одноразовый пароль")
			_, _ = server.Bot.Send(msg)
			return
		} else {
			msg := tgbotapi.NewMessage(chatID, "Неверный пароль")
			_, _ = server.Bot.Send(msg)
			return
		}
	} else if server.Db.CheckUser(chatID) && message == "/start" {
		msg := tgbotapi.NewMessage(chatID, "Добро пожаловать на курс")
		_, _ = server.Bot.Send(msg)
	}

	switch message {
	case "/generate secret_peter123":
		pass := generatePassword()
		server.Db.AddPassword(pass)

		msg := tgbotapi.NewMessage(chatID, pass)
		_, _ = server.Bot.Send(msg)
	case "Главное меню":
		msg := tgbotapi.NewMessage(chatID, "Выберите модуль")
		msg.ReplyMarkup = moduleKeys
		_, _ = server.Bot.Send(msg)

	default:
		module := server.Db.GetUserModule(chatID)
		HandleModule(chatID, module, message, server)
	}
}

func generatePassword() string {
	res, err := password.Generate(10, 3, 0, false, false)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func HandleKey(update tgbotapi.Update, bot *tgbotapi.BotAPI, db *database.DBAccess) {
	server := Server{Bot: bot, Db: db}

	chatID := update.CallbackQuery.Message.Chat.ID
	message := update.CallbackQuery.Data
	var parameter string

	if strings.Contains(message, " ") {
		array := strings.Split(message, " ")
		if array[0] == "/openModule" {
			message = array[0]
			parameter = array[1]
		}
	}

	switch message {
	case "/startCourse":
		server.Db.StartCourse(chatID)
		StartModule1(chatID, server)
	case "/nextModule":
		module := server.Db.GetUserModule(chatID)

		if !Answered(chatID, module, server) {
			return
		}

		server.Db.UpdateUserModule(chatID, module+1)
		HandleStart(chatID, module+1, server)
	case "/openModule":
		module, _ := strconv.ParseInt(parameter, 10, 64)
		HandleGet(chatID, module, server)
		//msg := tgbotapi.NewMessage(chatID, "Открыть модуль: "+parameter)
		//_, _ = server.Bot.Send(msg)
	}
}

func Answered(chatID, module int64, server Server) bool {

	ans1, ans2, ans3 := server.Db.GetTreeAnswers(chatID, module)
	if ans1 != "error" {
		if ans1 != "empty" && ans2 != "empty" && ans3 != "empty" {
			return true
		} else {
			return false
		}
	}

	ans1, ans2 = server.Db.GetTwoAnswers(chatID, module)
	if ans1 != "error" {
		if ans1 != "empty" && ans2 != "empty" {
			return true
		} else {
			return false
		}
	}

	ans1 = server.Db.GetAnswer(chatID, module)
	if ans1 != "empty" {
		return true
	}

	return false
}
