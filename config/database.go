package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDb() {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/laundry_db?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected")
	DB = db
}
