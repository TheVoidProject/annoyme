package main

import (
	"context"

	"github.com/peterbourgon/ff/v3/ffcli"
)

func NewReminder() *ffcli.Command {
	return &ffcli.Command{
		Name: "new",
		ShortUsage: "Create and Send a Notification",
		ShortHelp:  "Create and Send a Notification",
		// FlagSet:    nil,
		Exec: func(ctx context.Context, args []string) error {
			title := getInput("invalid input", "Title")
			message := getInput("invalid input", "Message")
			notify(title, message)
			return nil
		},
	}
}