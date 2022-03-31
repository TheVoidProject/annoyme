package main

import (
	"os"
	"log"

	n "github.com/martinlindhe/notify"
)

func notify(title string, body string) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("get pwd error")
		os.Exit(1)
	}
	n.Notify("AnnoyMe", title, body, pwd + "/assets/reminder.png")
}