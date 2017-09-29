package utils

import (
	"go.uber.org/zap"
)

var (
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

func InitLogger() {
	var err error
	Logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	Sugar = Logger.Sugar()
}
