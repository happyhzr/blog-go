package main

import (
	"github.com/Sirupsen/logrus"

	"github.com/insisthzr/blog-back/models"
	"github.com/insisthzr/blog-back/web"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Info("logger started")
	logrus.Info("postgres started")
	models.Start()
	logrus.Info("web starting")
	web.Start()
}
