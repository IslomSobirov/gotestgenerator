package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/testapp/helper"
)

//Connect to the db
func Connect() *sql.DB {
	enverr := godotenv.Load(".env")
	if enverr != nil {
		helper.LogError(enverr)
		log.Fatalf("Error loading .env file")
	}
	database := os.Getenv("DATABASE")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	db, err := sql.Open("mysql", user+":"+password+"@/"+database)

	// if there is an error opening the connection, handle it
	if err != nil {
		helper.LogError(err)
		log.Fatalf(err.Error())
	}

	return db
}
