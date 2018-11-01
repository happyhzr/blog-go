package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"

	"github.com/insisthzr/blog-go/tool"
)

var (
	db *gorm.DB
)

func Start() {
	var err error
	db, err = gorm.Open("mysql", viper.GetString("mysql.dsn"))
	tool.CheckError(err)

	GetDB().AutoMigrate(&User{}, &Post{}, &Category{}, &Tag{})

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)

	if viper.GetString("mode") != "prod" {
		db.LogMode(true)
	}
}

func GetDB() *gorm.DB {
	return db
}
