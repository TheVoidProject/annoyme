package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "print current annoyme version",
  Long:  `print the current version of annoyme`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("annoyme version - " + version)
  },
}