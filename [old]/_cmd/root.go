package cmd

import (
	"context"
	"fmt"
	"os"
	"flag"

	// "github.com/TheVoidProject/annoyme/pkg/notify"
	// "github.com/TheVoidProject/annoyme/pkg/prompt"
	flag "github.com/spf13/pflag"
	"github.com/peterbourgon/ff/v3/ffcli"
)



const VERSION = "0.0.1"

var (
	rootFlagSet = flag.NewFlagSet("annoyme", flag.ExitOnError)
	daemonFlag = rootFlagSet.String("daemon", "", "control the annoyme daemon options: install | uninstall | start | stop | restart | status")
)

var rootSubcommands = []*ffcli.Command{
	newCmd(),
}

var root  = &ffcli.Command{
		ShortUsage: "annoyme [flags] subcmd [flags] <required> [<optional> ...]",
		LongHelp: "Set reminders that persistently bug you as system messages until terminated or marked complete",
		UsageFunc: usage,
		FlagSet:	rootFlagSet,
		Exec: rootExec,
		Subcommands: rootSubcommands,
}

func rootExec(ctx context.Context, args []string) error {
	fmt.Println(args)
	return nil
}

func Execute() error {
	argc := len(os.Args[1:])

	err := root.Parse(os.Args[1:])
	if err != nil {
		return err
	}

	if argc < 1 {
		fmt.Println(root.UsageFunc(root))
		return nil
	} else {
		return root.Run(context.Background())
	}
}