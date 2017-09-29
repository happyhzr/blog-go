package config

import (
	"os"
)

type Config struct {
	Http Http `json:"http"`
	Db   Db   `json:"db"`
	Jwt  Jwt  `json:"jwt"`
}

type Http struct {
	Addr string `json:"addr"`
}

type Db struct {
	Dsn string `json:"dsn"`
}

type Jwt struct {
	Secret string `json:"secret"`
	Exp    int64  `json:"exp"`
}

var (
	DefaultConfig = Config{
		Http: Http{
			Addr: ":20001",
		},
		Db: Db{
			Dsn: "root:" + os.Getenv("MYSQL_PASS") + "@/blog?parseTime=true",
		},
		Jwt: Jwt{
			Secret: "keyboard cat",
			Exp:    24 * 60 * 60,
		},
	}
)
