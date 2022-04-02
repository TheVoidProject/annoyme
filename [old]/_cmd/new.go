package cmd

import (
	"context"

	"github.com/peterbourgon/ff/v3/ffcli"
)

func newCmd() *ffcli.Command {
	return &ffcli.Command{
		Name: "new",
		ShortUsage: "Create and Send a Notification",
		ShortHelp:  "Create and Send a Notification",
		// FlagSet:    nil,
		Exec: func(ctx context.Context, args []string) error {
			return nil
		},
	}
}
