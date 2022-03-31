package main

import (
	"time"

	"github.com/gen2brain/beeep"
)

func notify(title string, body string, delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
	err := beeep.Notify(title, body, "assets/reminder.png")
	if err != nil {
			panic(err)
	}
}