package model

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/insisthzr/blog-back/conf"
	"github.com/insisthzr/blog-back/db"
	"gopkg.in/mgo.v2"
)

type Archive struct {
	ID    bson.ObjectId `bson:"_id"`
	Year  int           `bson:"year"`
	Month int           `bson:"month"`
	Posts []*Post       `bson:"posts"`
}

func UpsertArchive(selector bson.M, update bson.M) (*mgo.ChangeInfo, error) {
	sess := db.CopySess()
	defer sess.Close()
	return sess.DB(conf.DBName).C(cArchive).Upsert(selector, update)
}

func ListArchives(query bson.M) ([]*Archive, error) {
	sess := db.CopySess()
	defer sess.Close()
	archives := []*Archive{}
	err := sess.DB(conf.DBName).C(cArchive).Find(query).Sort("-_id").All(&archives)
	return archives, err
}
