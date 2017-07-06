package db

import (
	"gopkg.in/mgo.v2"

	"github.com/insisthzr/blog-back/conf"
)

var (
	sess *mgo.Session
)

func Init() {
	var err error
	sess, err = mgo.Dial(conf.MongoURI)
	if err != nil {
		panic(err)
	}
	sess.SetMode(mgo.Monotonic, true)
}

func Ping() error {
	return sess.Ping()
}

func CopySess() *mgo.Session {
	return sess.Copy()
}
