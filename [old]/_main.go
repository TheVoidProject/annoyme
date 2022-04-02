package main

import (
	"log"

	"github.com/TheVoidProject/annoyme/cmd"
)

const VERSION = "0.0.1"

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}