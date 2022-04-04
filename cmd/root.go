/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/TheVoidProject/annoyme/pkg/daemon"
	"github.com/TheVoidProject/annoyme/pkg/logger"

)

// var log = logrus.New()
// var stdout = logrus.New()
var (
	stdout logrus.Logger
	log logrus.Logger
)
var cfgFile string
var daemonFlag string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "annoyme",
	Short: "annoyme [flags] subcmd [flags] <required> [<optional> ...]",
	Long: "Set reminders that persistently bug you as system messages until terminated or marked complete",
	Run: rootRun,
}

func rootRun(cmd *cobra.Command, args []string) {
	if len(daemonFlag) > 0 {
		daemon.Control(daemonFlag)
	} else {
		cmd.Usage()
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetUsageTemplate(usageTemplate())
	stdout, log = logger.New("cmd")
	// cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.annoyme.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&daemonFlag, "daemon", "d", "", "annoyme [--daemon|-d] install | uninstall | start | stop | status")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".annoyme" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".annoyme")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// func initLoggers() {
// 	currentTime := time.Now()
// 	logPath := filepath.Join("logs", currentTime.Format("2006-01-02_150405.log"))
// 	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	if err == nil {
// 		log.Out = file
// 	} else {
// 		log.Info("Failed to log to file, using default stderr")
// 	}

// 	stdout.Formatter = new(logrus.TextFormatter)
// 	log.Formatter = new(logrus.JSONFormatter) // json
// 	// log.Formatter.(*logrus.TextFormatter).DisableColors = true    // remove colors
// 	// log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output
// 	log.Level = logrus.InfoLevel // default
// 	stdout.Level = logrus.InfoLevel
// 	log.SetReportCaller(true)    // shows where is was called from
// 	stdout.SetReportCaller(true) // shows where is was called from
// }