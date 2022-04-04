package logger

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()
var stdout = logrus.New()


func New(filename string) (logrus.Logger, logrus.Logger) {
	logPath := filepath.Join("logs", filename + ".log")
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	stdout.Formatter = new(logrus.TextFormatter)
	log.Formatter = new(logrus.JSONFormatter) // json
	// log.Formatter.(*logrus.TextFormatter).DisableColors = true    // remove colors
	// log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	log.Level = logrus.InfoLevel // default
	stdout.Level = logrus.InfoLevel
	log.SetReportCaller(true)    // shows where is was called from
	stdout.SetReportCaller(true) // shows where is was called from
	return *stdout, *log
}