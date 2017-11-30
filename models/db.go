package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/insisthzr/blog-back/config"
)

var (
	db *sqlx.DB
)

func DB() *sqlx.DB {
	return db
}

func Start() {
	db = sqlx.MustConnect("mysql", config.C.Db.Dsn)
}
