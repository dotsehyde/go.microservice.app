package logger

import (
	"os"

	"go.uber.org/zap"
)

var Log = func() *zap.Logger {
	var logger *zap.Logger
	var err error
	if os.Getenv("APP_ENV") == "production" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		panic(err)
	}
	return logger
}()
