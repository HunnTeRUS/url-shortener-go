package logger

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestInfo(t *testing.T) {
	infoMessage := "logging info message"
	tags := zap.Field{
		Key:    "anotherMessage-Key",
		String: "anotherMessage-Value",
		Type:   zapcore.StringType,
	}

	Info(infoMessage, tags)
}

func TestError(t *testing.T) {
	infoMessage := "logging error message"
	err := errors.New("logging an error message")
	tags := zap.Field{
		Key:    "anotherMessage-Key",
		String: "anotherMessage-Value",
		Type:   zapcore.StringType,
	}

	Error(infoMessage, err, tags)
}

func TestGetLogger(t *testing.T) {
	log := GetLogger()
	assert.NotNil(t, log)
}

func TestGetOutput(t *testing.T) {
	os.Setenv(envLogOutput, "stderr")
	out := getOutput()
	assert.Equal(t, out, "stderr")
}

func TestGetLevel(t *testing.T) {
	os.Setenv(envLogLevel, "info")
	logLevel := getLevel()
	assert.Equal(t, logLevel, zap.InfoLevel)

	os.Setenv(envLogLevel, "error")
	logLevel = getLevel()
	assert.Equal(t, logLevel, zap.ErrorLevel)

	os.Setenv(envLogLevel, "debug")
	logLevel = getLevel()
	assert.Equal(t, logLevel, zap.DebugLevel)

	os.Setenv(envLogLevel, "")
	logLevel = getLevel()
	assert.Equal(t, logLevel, zap.InfoLevel)
}

func TestPrint(t *testing.T) {
	infoMessage := "logging message to specific print method"
	logger := GetLogger()
	logger.Print(infoMessage, infoMessage)

	logger = GetLogger()
	logger.Print()
}

func TestPrintf(t *testing.T) {
	infoMessage := "logging message to specific print method"
	logger := GetLogger()
	logger.Printf("test", infoMessage, infoMessage)

	logger.Printf("test")
}
