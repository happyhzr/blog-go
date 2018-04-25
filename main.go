package main

import (
	"github.com/insisthzr/blog-back/api"
	"github.com/insisthzr/blog-back/model"
	"github.com/insisthzr/blog-back/tool/logger"
)

func main() {
	logger.Init()
	model.Start()
	api.Start()
}
