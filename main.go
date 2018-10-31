package main

import (
	"github.com/insisthzr/blog-go/api"
	"github.com/insisthzr/blog-go/model"
	"github.com/insisthzr/blog-go/tool/logger"
)

func main() {
	logger.Init()
	model.Start()
	api.Start()
}
