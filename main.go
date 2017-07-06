package main

import (
	"github.com/insisthzr/blog-back/conf"
	"github.com/insisthzr/blog-back/db"
	"github.com/insisthzr/blog-back/route"
)

func main() {
	db.Init()
	e := route.Init()
	e.Start(":" + conf.HTTPPort)
}
