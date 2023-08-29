package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)
	log.SetOutput(os.Stdout)
}

func GetLogger() *logrus.Logger {
	return log
}
