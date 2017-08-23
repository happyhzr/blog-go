package main

import (
	"github.com/insisthzr/blog-back/config"
	"github.com/insisthzr/blog-back/router"
	"github.com/insisthzr/blog-back/service/mysql"
)

func main() {
	config.Load()
	mysql.Run()
	router.Run()
}
