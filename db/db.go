package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
	db, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Cant open db")
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	DB = db
	createTables()

}

func createTables() {
	createEventTable := `
	  CREATE TABLE IF NOT EXISTS events (
	  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	  name TEXT NOT NULL,
	  description TEXT NOT NULL,
	  location TEXT NOT NULL,
	  dateTime DATETIME NOT NULL,
	  user_id INTEGER,
	  FOREIGN KEY(user_id) REFERENCES users(id)
  	);`

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	username TEXT,
	password TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE);
`

	_, err := DB.Exec(createEventTable)
	if err != nil {
		fmt.Printf("Could not create Events table %v", err)
	}
	fmt.Println("Events table create")

	_, err = DB.Exec(createUsersTable)
	if err != nil {
		fmt.Printf("Could not create Users table %v", err)
	}
	fmt.Println("Users table create")
}
