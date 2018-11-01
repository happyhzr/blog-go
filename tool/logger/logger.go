package logger

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	Sugar *zap.SugaredLogger
)

func Init() {
	var logger *zap.Logger
	var err error
	if viper.GetString("mode") == "prod" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		panic(err)
	}
	Sugar = logger.Sugar()
}
