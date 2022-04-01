package main

import (
	"time"

	"github.com/gen2brain/beeep"
)


		// "idReminder" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		// "title"       TEXT,
		// "message"     TEXT,
		// "date"        INTEGER,
		// "time"        INTEGER,
		// "repeat"      INTEGER,
		// "repeatDelay" INTEGER,

type notification struct {
	title string
	message string
	datetime time.Time
	repeat int
	delay int
}


func (n notification) notify() {
	time.Sleep(time.Duration(n.delay) * time.Second)
	err := beeep.Notify(n.title, n.message, "assets/reminder.png")
	if err != nil {
			panic(err)
	}
}