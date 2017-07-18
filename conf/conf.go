package conf

var (
	HTTPPort = "1324"

	JWTKey = []byte("keyboard cat")
	JWTExp = 24 * 60 * 60

	MongoURI = "localhost"
	DBName   = "blog"

	FilePath = "/var/www/blog/file/"
	LogPath  = "/mnt/log/blog"
)
