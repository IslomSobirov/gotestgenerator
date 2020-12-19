package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	database = "user_db"
	user     = "sammy"
	password = "password"
)

//Connect to the db
func Connect() *sql.DB {
	db, err := sql.Open("mysql", user+":"+password+"@/"+database)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	return db
}
