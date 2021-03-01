package logs

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"os"
)

type Logger struct {
	jsonLogger log.Logger
}

func NewLogger() *Logger {
	jsonLogger := log.NewJSONLogger(os.Stdout)
	jsonLogger = log.With(jsonLogger, "ts", log.DefaultTimestampUTC)
	jsonLogger = log.With(jsonLogger, "caller", log.DefaultCaller)

	return &Logger{
		jsonLogger: jsonLogger,
	}
}

func (l Logger) Info(message string) {
	logErr := l.jsonLogger.Log("type", "info", "message", message)
	if logErr != nil {
		fmt.Println(logErr, message)
	}
}

func (l Logger) Error(err error)  {
	logErr := l.jsonLogger.Log("type", "error", "message", err.Error())
	if logErr != nil {
		fmt.Println(logErr, err)
	}
}


func (l Logger) Fatal(err error)  {
	logErr := l.jsonLogger.Log("type", "fatal", "message", err.Error())
	if logErr != nil {
		fmt.Println(logErr, err)
	}

	os.Exit(1)
}