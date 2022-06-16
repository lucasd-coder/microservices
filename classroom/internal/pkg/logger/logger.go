package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var Log = log.WithFields(log.Fields{
	"logName":  "classroom",
	"logIndex": "message",
})

const defaultLogLevel = "debug"

func SetUpLog() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	logEnvLevel := os.Getenv("LOG_LEVEL")
	if logEnvLevel == "" {
		logEnvLevel = defaultLogLevel
	}

	logLevel, _ := log.ParseLevel(logEnvLevel)
	log.SetLevel(logLevel)
}
