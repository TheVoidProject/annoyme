package main

import (
	"os"
	"fmt"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func openDatabase() error {
	var err error
	// temporary we will write the db to ~/.config/annoyme/* for unix and AppData/roaming/annoyme/* for windows
	db, err = sql.Open("sqlite3", "./bin/annoyme-sqlite.db")
	if err != nil {
		return err
	}

	return db.Ping()
}

func createTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS reminder (
		"idReminder" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"title"       TEXT,
		"message"     TEXT,
		"date"        INTEGER,
		"time"        INTEGER,
		"repeat"      INTEGER,
		"repeatDelay" INTEGER,
	  );`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, execErr := statement.Exec()
	if execErr != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	log.Println("AnnoyMe Reminder table created")
}