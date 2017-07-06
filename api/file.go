package api

import (
	"io"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	"github.com/insisthzr/blog-back/busniess"
	"github.com/insisthzr/blog-back/conf"
	"github.com/insisthzr/blog-back/util"
)

func UploadFile(c echo.Context) error {
	userID := util.GetIDFromJWT(c.Get("user").(*jwt.Token))
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(400, JSON{"msg": err.Error()})
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(400, JSON{"msg": err.Error()})
	}
	defer src.Close()

	// Destination
	path := conf.FilePath + file.Filename
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()
	// Copy
	written, err := io.Copy(dst, src)
	if err != nil {
		return err
	}
	in := &busniess.CreateFileIn{
		FileIn: busniess.FileIn{
			Name:      file.Filename,
			Path:      path,
			Size:      written,
			CreatedBy: userID,
		},
	}
	out := busniess.CreateFile(in)
	return c.JSON(200, JSON{"file": out})
}

func ListFiles(c echo.Context) error {
	out := busniess.ListFiles()
	return c.JSON(200, JSON{"files": out})
}
