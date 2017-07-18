package api

import (
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"

	"github.com/insisthzr/blog-back/busniess"
)

func ListArchives(c echo.Context) error {
	out := busniess.ListArchives()
	logrus.WithFields(logrus.Fields{"out": out}).Infoln("ListArchives")
	return c.JSON(200, JSON{"archives": out})
}
