package utils

import "go.uber.org/zap"

var Logger *zap.Logger

func InitializeLogger() {
	Logger, _ = zap.NewProduction()
}

func Loggers() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
