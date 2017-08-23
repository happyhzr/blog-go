package config

import (
	"os"
)

var (
	HttpAddr = "localhost:20000"

	MysqlDsn = "root:root@/blog"
)

func Load() {
	setXX(&HttpAddr, "HTTP_ADDR")
	setXX(&MysqlDsn, "MYSQL_DSN")
}

func setXX(value *string, key string) {
	env := os.Getenv(key)
	if env != "" {
		*value = env
	}
}
