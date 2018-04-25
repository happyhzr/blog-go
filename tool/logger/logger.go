package logger

import (
	"go.uber.org/zap"
)

var (
	Sugar *zap.SugaredLogger
)

func Init() {
	var err error
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	Sugar = logger.Sugar()
}
