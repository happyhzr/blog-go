package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/insisthzr/blog-back/config"
	"github.com/insisthzr/blog-back/util"
)

var (
	db *sql.DB
)

func Run() {
	var err error
	db, err = sql.Open("mysql", config.MysqlDsn)
	util.CheckError(err)
}

func GetDb() *sql.DB {
	return db
}
