package api

import (
	"github.com/labstack/echo"

	"github.com/insisthzr/blog-back/busniess"
)

func ListArchives(c echo.Context) error {
	archives := busniess.ListArchives()
	return c.JSON(200, JSON{
		"archives": archives,
	})
}
