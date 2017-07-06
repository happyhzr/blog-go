package model

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/insisthzr/blog-back/conf"
	"github.com/insisthzr/blog-back/db"
)

type File struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string        `bson:"name"`
	Path      string        `bson:"path"`
	Size      int64         `bson:"int"`
	CreatedBy bson.ObjectId `bson:"created_by,omitempty"`
	CreatedAt int64         `bson:"created_at"`
}

func (f *File) Save() error {
	sess := db.CopySess()
	defer sess.Close()
	f.ID = bson.NewObjectId()
	return sess.DB(conf.DBName).C(cFile).Insert(f)
}

func ListFiles(query bson.M) ([]*File, error) {
	sess := db.CopySess()
	defer sess.Close()
	files := []*File{}
	err := sess.DB(conf.DBName).C(cFile).Find(query).All(&files)
	return files, err
}

func GetFile(query bson.M) (*File, error) {
	sess := db.CopySess()
	defer sess.Close()
	file := &File{}
	err := sess.DB(conf.DBName).C(cFile).Find(query).One(&file)
	return file, err
}
