package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/insisthzr/blog-back/utils"
)

type Config struct {
	Debug bool `json:"debug"`
	Http  Http `json:"http"`
	Db    Db   `json:"db"`
	Jwt   Jwt  `json:"jwt"`
}

type Http struct {
	Addr string `json:"addr"`
}

type Db struct {
	Dsn string `json:"dsn"`
}

type Jwt struct {
	Secret string `json:"secret"`
	Maxage int64  `json:"maxage"`
}

var C *Config

func Load(path string) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	C = &Config{}
	err = json.Unmarshal(b, C)
	if err != nil {
		panic(err)
	}
	C.Db.Dsn = strings.Replace(C.Db.Dsn, "{{password}}", os.Getenv("MYSQL_PASS"), -1)
	utils.Sugar.Infow("load config", "config", C)
}
