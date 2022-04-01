package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/peterbourgon/ff/v3/ffcli"
)

const VERSION = "0.0.1"

var (
	rootFlagSet   = flag.NewFlagSet("annoyme", flag.ExitOnError)
	verboseFlag   = rootFlagSet.Bool("v", false, "increase log verbosity")
	
	daemonFlag		= rootFlagSet.String("daemon", "", "control the daemon process of annoyme")
)
// var (
// 	daemonFlagSet = flag.NewFlagSet("daemon", flag.ExitOnError)
// 	svcFlag 			= daemonFlagSet.String("c", "", "Control the system service. options: install, uninstall, start, stop, restart")
// )


func main() {
	flag.Parsed()
	root := &ffcli.Command{
		ShortUsage: "annoyme [flags] subcmd [flags] <required> [<optional> ...]",
		LongHelp: "Set reminders that persistently bug you as system messages until terminated or marked complete",
		UsageFunc: customUsage,
		FlagSet:     rootFlagSet,
		Exec: func(ctx context.Context, args []string) error {
			if *verboseFlag == true {
				print("in verbose")
			}
			if len(*daemonFlag) != 0 {
				daemon(*daemonFlag)
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