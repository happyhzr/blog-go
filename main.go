package main

import (
	"github.com/Sirupsen/logrus"

	"github.com/insisthzr/blog-back/conf"
	"github.com/insisthzr/blog-back/db"
	"github.com/insisthzr/blog-back/route"
	"github.com/insisthzr/blog-back/util"
)

func main() {
	go util.AutoRotateLog()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	db.Init()
	e := route.Init()
	e.Start(":" + conf.HTTPPort)
}
