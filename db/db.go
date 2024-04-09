package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		panic(err)
	}
	DB = db;
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	fmt.Println(DB)
	createTables()
}

func createTables() {
	sql := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`
	_, err := DB.Exec(sql)
	if err != nil {
		panic(err)
	}

	// create events table
	sql = `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT, description TEXT NOT NULL,
		location TEXT NOT NULL, 
		datetime DATETIME NOT NULL, 
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
		)`
	_, err = DB.Exec(sql)
	if err != nil {
		panic(err)
	}
}
