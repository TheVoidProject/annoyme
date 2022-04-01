package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"


	"github.com/peterbourgon/ff/v3/ffcli"
)

const VERSION = "0.0.1"

var (
	rootFlagSet   = flag.NewFlagSet("annoyme", flag.ExitOnError)
	verboseFlag   = rootFlagSet.Bool("v", false, "increase log verbosity")
	svcFlag    = rootFlagSet.String("daemon", "", "the help string")
)


func main() {
	root := &ffcli.Command{
		ShortUsage: "annoyme [flags] subcmd [flags] <required> [<optional> ...]",
		LongHelp: "Set reminders that persistently bug you as system messages until terminated or marked complete",
		UsageFunc: customUsage,
		FlagSet:     rootFlagSet,
		Exec: func(ctx context.Context, args []string) error {
			fmt.Println("main")
			if len(*svcFlag) != 0 {
				daemon()
			}
			// if *verboseFlag == true {
			// 	print("in verbose")
			// }
			// fmt.Println("args: "+ strings.Join(args, ","))
			return nil
		},
		Subcommands: []*ffcli.Command{},
	}
	if err := root.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		flag.Parse()
		log.Fatal(err)
	}
}