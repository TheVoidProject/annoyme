package data

import (
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() error {
	var err error
	// temporary we will write the db to ~/.config/annoyme/* for unix and AppData/roaming/annoyme/* for windows
	db, err = sql.Open("sqlite3", "./annoyme-sqlite.db")
	if err != nil {
		return err
	}

	return db.Ping()
}

func CreateTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS reminder (
		"idReminder" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"title" TEXT,
		"message" TEXT,
		"level" TEXT
	  );`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("AnnoyMe Reminder table created")
}