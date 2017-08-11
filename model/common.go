package model

const (
	cUser    = "users"
	cPost    = "posts"
	cArchive = "archives"
	cFile    = "files"
)

type QueryConfig struct {
	Skip  int
	Limit int
}
