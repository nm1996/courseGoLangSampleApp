package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

const (
	jwtSecretKey = "your_secret_key" // Replace with your own secret key
)

func initDB() {
	var err error
	// Open a connection to the database
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/sample_nina_db")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
}

func GetDBConnection() *sql.DB {
	initDB()
	return db
}
