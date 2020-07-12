package logger

import (
	"os"

	"go.uber.org/zap"
)

var logger, _ = zap.NewProduction(zap.AddCallerSkip(1))
var sugar = logger.Sugar()

func Warning(message interface{}) {
	sugar.Warn(message)
}

func Error(message interface{}) {
	sugar.Error(message)
	os.Exit(1)
}

func Info(message interface{}) {
	sugar.Info(message)
}
