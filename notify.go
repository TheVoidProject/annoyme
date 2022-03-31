package main

import (
	"github.com/gen2brain/beeep"
)

func notify(title string, body string) {
	err := beeep.Notify(title, body, "assets/reminder.png")
	if err != nil {
			panic(err)
	}
}