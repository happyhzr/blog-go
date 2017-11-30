package main

import (
	"github.com/insisthzr/blog-back/config"
	"github.com/insisthzr/blog-back/models"
	"github.com/insisthzr/blog-back/utils"
	"github.com/insisthzr/blog-back/web"
)

func main() {
	utils.InitLogger()
	config.Load("/mnt/code/blog/config.json")
	models.Start()
	web.Start()
}
