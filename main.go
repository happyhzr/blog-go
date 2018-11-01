package main

import (
	"github.com/spf13/viper"

	"github.com/insisthzr/blog-go/api"
	"github.com/insisthzr/blog-go/model"
	"github.com/insisthzr/blog-go/tool"
	"github.com/insisthzr/blog-go/tool/logger"
)

func main() {
	logger.Init()
	logger.Sugar.Infow("logger", "start", "OK")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	tool.CheckError(err)
	logger.Sugar.Infow("model", "start", "OK")
	model.Start()
	logger.Sugar.Infow("model", "start", "OK")
	api.Start()
}
