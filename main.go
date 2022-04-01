package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	// "time"
	"github.com/peterbourgon/ff/v3/ffcli"
)

const VERSION = "0.0.1"

var (
	rootFlagSet   = flag.NewFlagSet("annoyme", flag.ExitOnError)
	verbose       = rootFlagSet.Bool("v", false, "increase log verbosity")
)


func main() {
	root := &ffcli.Command{
		ShortUsage: "annoyme [flags] subcmd [flags] <required> [<optional> ...]",
		LongHelp: "Set reminders that persistently bug you as system messages until terminated or marked complete",
		UsageFunc: usage,
		FlagSet:     rootFlagSet,
		Exec: func(ctx context.Context, args []string) error {
			if *verbose == true {
				print("in verbose")
			}
			fmt.Println("args: "+ strings.Join(args, ","))
			return nil
		},
		Subcommands: []*ffcli.Command{},
	}
	if err := root.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}