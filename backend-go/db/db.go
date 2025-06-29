package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

// createTables creates the necessary tables in the database if they do not exist.
func createTables() {
	createUsersTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := DB.Exec(createUsersTableSQL)
	if err != nil {
		panic("Could not create users table: " + err.Error())
	}

	createEventsTableSQL := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	_, err = DB.Exec(createEventsTableSQL)
	if err != nil {
		panic("Could not create events tables: " + err.Error())
	}

	createRegistrationsTableSQL := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY (event_id) REFERENCES events(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	_, err = DB.Exec(createRegistrationsTableSQL)
	if err != nil {
		panic("Could not create registrations table: " + err.Error())
	}
}
