package logs

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"os"
)

type Logger interface {
	Info(message string)
	Error(err error)
	Fatal(err error)
}

type JSONLogger struct {
	jsonLogger log.Logger
}

func NewLogger() Logger {
	jsonLogger := log.NewJSONLogger(os.Stdout)
	jsonLogger = log.With(jsonLogger, "ts", log.DefaultTimestampUTC)
	jsonLogger = log.With(jsonLogger, "caller", log.Caller(4))

	return &JSONLogger{
		jsonLogger: jsonLogger,
	}
}

func (l JSONLogger) Info(message string) {
	logErr := l.jsonLogger.Log("type", "info", "message", message)
	if logErr != nil {
		fmt.Println(logErr, message)
	}
}

func (l JSONLogger) Error(err error) {
	logErr := l.jsonLogger.Log("type", "error", "message", err.Error())
	if logErr != nil {
		fmt.Println(logErr, err)
	}
}

func (l JSONLogger) Fatal(err error) {
	logErr := l.jsonLogger.Log("type", "fatal", "message", err.Error())
	if logErr != nil {
		fmt.Println(logErr, err)
	}

	os.Exit(1)
}
