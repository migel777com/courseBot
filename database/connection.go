package database

import (
	"courseBot/logger"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

type DBAccess struct {
	Db *sql.DB
	logger *logger.Logger
}

func OpenDB(logger *logger.Logger) (DBAccess, error) {

	_ = godotenv.Load("globals.env")
	db_username := os.Getenv("db_username")
	db_password := os.Getenv("db_password")
	server_address := os.Getenv("server_address")
	dp_port := os.Getenv("db_port")
	db_name := os.Getenv("db_name")
	db, err := sql.Open("mysql", db_username+":"+db_password+"@tcp("+server_address+":"+dp_port+")/"+db_name)

	if err != nil {
		logger.MakeLog("Error occurred while setting the connection with DB"+err.Error())
		panic(err.Error())
	}

	res := DBAccess{db, logger}
	return res, err
}
