package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func DbConnection() {
	var err error
	connectionString := "root:Richard658958!@tcp(127.0.0.1:3306)/GoProject"
	DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("DB ping error:", err)
	}

	log.Println("MySQL connected.")
}
