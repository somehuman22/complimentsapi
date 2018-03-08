package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

//DB ... our global database
var DB *sql.DB
var err error

//InitDB ... assigns database
func InitDB() {
	DB, err = sql.Open("sqlite3", "dev.db")
}
