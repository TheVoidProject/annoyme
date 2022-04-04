/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	// "fmt"
	"log"
	"net"
	// "strings"
	"time"

	// "encoding/gob"

	"github.com/TheVoidProject/annoyme/pkg/prompt"
	"github.com/TheVoidProject/annoyme/pkg/reminder"
	"github.com/spf13/cobra"
)

func newRun(cmd *cobra.Command, args []string) {
	conn, err := net.Dial("tcp", ":9977")
	if err != nil {
			log.Fatal("Connection error", err)
	}
	title := prompt.GetString("Invalid Format", "Title")
	msg := prompt.GetString("Invalid Format", "Message")
	r := reminder.New(title, msg)
	t := prompt.GetTime("Invalid Time Format")
	r.Time = t
	sound := prompt.GetBool("Invalid Format must be y|n", "Should I yell?")
	r.Sound = sound
	shouldNag := prompt.GetBool("Invalid Format must be y|n", "Do you want me to nag you?")
	if shouldNag {
		repeat := prompt.GetInt("Invalid format must be an integer 1...9+", "Number of nags")
		r.Repeat = repeat
		interval := prompt.GetInt("Invalid format must be an integer 1...9+", "Nagging interval? [minutes]")
		r.Delay = time.Duration(interval) * time.Minute
	}
	reminder.Encode(r, conn)
	conn.Close()
}

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: newRun,
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
