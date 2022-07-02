package database

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (db *DBAccess) AddNewUser(chatId int64, firstName, lastName, tgUsername string) {
	stmt := "INSERT INTO users(chat_id, first_name, second_name, telegram_username) VALUES (?, ?, ?, ?);"
	db.Db.Exec(stmt, chatId, firstName, lastName, tgUsername)
}

func (db *DBAccess) CheckUser(chatID int64) bool {
	stmt := "SELECT * FROM users WHERE chat_id = ?;"
	results, err := db.Db.Query(stmt, chatID)

	defer results.Close()

	if err != nil {
		log.Println("Error occurred while checking user existence from database ->", err)
		db.logger.MakeLog("Error occurred while checking user existence from database ->" + err.Error())
	}
	if results.Next() {
		return true
	}
	return false
}

func (db *DBAccess) AddPassword(password string) {
	stmt := "INSERT INTO password(password) VALUES (?);"
	db.Db.Exec(stmt, password)
}

func (db *DBAccess) CheckPassword(update tgbotapi.Update) bool {
	pass := update.Message.Text

	stmt := "SELECT status FROM password WHERE password = ?;"
	results, err := db.Db.Query(stmt, pass)
	var status string
	defer results.Close()

	if err != nil {
		log.Println("Error occurred while checking password ->", err)
		db.logger.MakeLog("Error occurred while checking password ->" + err.Error())
	}
	if results.Next() {
		results.Scan(&status)

		if status == "generated" {
			query := "UPDATE password SET status = 'activated' WHERE password = ?;"
			db.Db.Exec(query, pass)

			db.AddNewUser(update.Message.Chat.ID, update.Message.From.FirstName, update.Message.From.LastName, update.Message.From.UserName)
		}

		return true
	}
	return false
}

func (db *DBAccess) StartCourse(chatID int64) {
	stmt := "INSERT INTO answers(chat_id) VALUES (?);"
	db.Db.Exec(stmt, chatID)

	query := "UPDATE users SET module = 1 WHERE chat_id = ?;"
	db.Db.Exec(query, chatID)
}

func (db *DBAccess) GetUserModule(chatID int64) int64 {
	stmt := "SELECT module FROM users WHERE chat_id = ?;"
	results, err := db.Db.Query(stmt, chatID)
	var module int64
	defer results.Close()

	if err != nil {
		log.Println("Error occurred while getting user module ->", err)
	}

	if results.Next() {
		results.Scan(&module)
	}

	return module
}

func (db *DBAccess) UpdateUserModule(chatID, module int64) {
	stmt := "UPDATE users SET module = ? WHERE chat_id = ?;"
	db.Db.Exec(stmt, module, chatID)
}

func (db *DBAccess) GetAnswer(chatID int64, module int64) string {
	question := "ans" + strconv.FormatInt(module, 10) + "_1"

	stmt := "SELECT " + question + " FROM answers WHERE chat_id = ?;"
	results, err := db.Db.Query(stmt, chatID)
	if err != nil {
		//log.Println("Error occurred while getting user answer ->", err)
		return "no questions"
	}

	defer results.Close()

	if results.Next() {
		results.Scan(&question)
	}

	if question == "ans"+strconv.FormatInt(module, 10)+"_1" {
		question = "empty"
	}

	return question
}

func (db *DBAccess) GetTwoAnswers(chatID int64, module int64) (string, string) {
	question1 := "ans" + strconv.FormatInt(module, 10) + "_1"
	question2 := "ans" + strconv.FormatInt(module, 10) + "_2"

	stmt := "SELECT " + question1 + "," + question2 + " FROM answers WHERE chat_id = ?;"
	results, err := db.Db.Query(stmt, chatID)
	if err != nil {
		//log.Println("Error occurred while getting user answers ->", err)
		return "error", ""
	}

	defer results.Close()

	if results.Next() {
		results.Scan(&question1, &question2)
	}

	if question1 == "ans"+strconv.FormatInt(module, 10)+"_1" {
		question1 = "empty"
	}

	if question2 == "ans"+strconv.FormatInt(module, 10)+"_2" {
		question2 = "empty"
	}

	return question1, question2
}

func (db *DBAccess) GetTreeAnswers(chatID int64, module int64) (string, string, string) {
	question1 := "ans" + strconv.FormatInt(module, 10) + "_1"
	question2 := "ans" + strconv.FormatInt(module, 10) + "_2"
	question3 := "ans" + strconv.FormatInt(module, 10) + "_3"

	stmt := "SELECT " + question1 + "," + question2 + "," + question3 + " FROM answers WHERE chat_id = ?;"
	results, err := db.Db.Query(stmt, chatID)
	if err != nil {
		//log.Println("Error occurred while getting user answers ->", err)
		return "error", "", ""
	}

	defer results.Close()

	if results.Next() {
		results.Scan(&question1, &question2, &question3)
	}

	if question1 == "ans"+strconv.FormatInt(module, 10)+"_1" {
		question1 = "empty"
	}

	if question2 == "ans"+strconv.FormatInt(module, 10)+"_2" {
		question2 = "empty"
	}

	if question3 == "ans"+strconv.FormatInt(module, 10)+"_3" {
		question3 = "empty"
	}

	return question1, question2, question3
}

func (db *DBAccess) AnswerQuestion(chatID int64, question, answer string) {
	stmt := "UPDATE answers SET " + question + "=? WHERE chat_id = ?;"
	db.Db.Exec(stmt, answer, chatID)
}
