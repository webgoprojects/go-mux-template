package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func Init(logLevel string) error {
	var err error

	// Use development config for development/debug environment, production otherwise
	if logLevel == "debug" || logLevel == "development" {
		Logger, err = zap.NewDevelopment()
	} else {
		Logger, err = zap.NewProduction()
	}

	if err != nil {
		return err
	}
	return nil
}

func Sync() {
	if Logger != nil {
		Logger.Sync()
	}
}
