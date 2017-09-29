package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/insisthzr/blog-back/config"
)

var (
	db *sql.DB
)

func DB() *sql.DB {
	return db
}

func Start() {
	var err error
	db, err = sql.Open("mysql", config.DefaultConfig.Db.Dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
