package models

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/insisthzr/blog-back/config"
)

var (
	db *sql.DB
)

func getDb() *sql.DB {
	return db
}

func Start() {
	var err error
	db, err = sql.Open("postgres", config.DefaultConfig.Db.Dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
