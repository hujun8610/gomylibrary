package logger

import "testing"

func TestGetLoggerWithFile(t *testing.T) {
	log := GetLoggerWithFile("test.log")
	log.Info("test")
	log.Error("test")
}
