package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func ConnectDB() {
	db, err = sql.Open("mysql", "root:Biplove@123@tcp(127.0.0.1:3306)/go_tododb")
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return db
}
