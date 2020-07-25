package logger

import (
	"Sparkle/config"

	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	l, _ := zap.NewDevelopment(zap.AddCallerSkip(1))
	logger = l.Sugar()
}

func Setup(config config.LoggerConfig) error {

	var lc zap.Config

	if config.Debug {
		lc = newDebugConfig()
	} else {
		lc = newProductionConfig()
	}

	l, err := lc.Build(zap.AddCallerSkip(1))

	if err != nil {
		return err
	}

	logger = l.Sugar()

	return nil
}

func newProductionConfig() zap.Config {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	return config
}

func newDebugConfig() zap.Config {
	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	return config
}

func Debug(message interface{}) {
	logger.Debug(message)
}

func Debugf(message string, args ...interface{}) {
	logger.Debugf(message, args)
}

func Warning(message interface{}) {
	logger.Warn(message)
}

func Warningf(message string, args ...interface{}) {
	logger.Warnf(message, args...)
}

func Error(message interface{}) {
	logger.Error(message)
}

func Errorf(message string, args ...interface{}) {
	logger.Errorf(message, args)
}

func Fatal(message interface{}) {
	logger.Fatal(message)
}

func Fatalf(message string, args ...interface{}) {
	logger.Fatalf(message, args)
}
