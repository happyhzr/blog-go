package config

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
			Dsn: "user=root port=5433 dbname=blog",
		},
		Jwt: Jwt{
			Secret: "keyboard cat",
			Exp:    24 * 60 * 60,
		},
	}
)
