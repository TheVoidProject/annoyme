package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/peterbourgon/ff/v3/ffcli"
)

const VERSION = "0.0.1"


func main() {
	fs := flag.NewFlagSet("notify", flag.ExitOnError)
	// t := fs.String("t", "title", "title of the notification")
	root := &ffcli.Command{
		ShortUsage: "Sends a Notification",
		ShortHelp:  "Sends a Notification",
		FlagSet:    fs,
		Exec: func(ctx context.Context, args []string) error {
			return nil
		},
		Subcommands: []*ffcli.Command{newReminder()},
	}

	if err := root.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}