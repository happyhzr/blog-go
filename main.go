package main

import (
	"github.com/insisthzr/blog-back/models"
	"github.com/insisthzr/blog-back/utils"
	"github.com/insisthzr/blog-back/web"
)

func main() {
	utils.InitLogger()
	utils.Sugar.Info("logger started")
	models.Start()
	utils.Sugar.Info("mysql started")
	utils.Sugar.Info("web starting")
	web.Start()
}
