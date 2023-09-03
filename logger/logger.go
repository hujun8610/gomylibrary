package logger

import (
	"github.com/sirupsen/logrus"
	"io"
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

func GetLoggerWithFile(filename string) *logrus.Logger {
	file, err := os.OpenFile(filename+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return log
	}
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
	return log
}
