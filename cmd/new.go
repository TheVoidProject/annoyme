/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net"
	// "encoding/gob"

	"github.com/spf13/cobra"
	"github.com/TheVoidProject/annoyme/pkg/reminder"
	"github.com/TheVoidProject/annoyme/pkg/prompt"
)

func newRun(cmd *cobra.Command, args []string) {
	conn, err := net.Dial("tcp", ":9977")
	if err != nil {
			log.Fatal("Connection error", err)
	}
	title := prompt.GetString("invalid format", "Title")
	msg := prompt.GetString("invalid format", "Message")
	r := reminder.Reminder{
		Title: title,
		Message: msg,
	}
	reminder.Encode(r, conn)
	conn.Close()
	fmt.Println("new called")
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
