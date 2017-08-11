package model

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/insisthzr/blog-back/conf"
	"github.com/insisthzr/blog-back/db"
)

type Post struct {
	ID        bson.ObjectId `bson:"_id"`
	Title     string        `bson:"title"`
	Body      string        `bson:"body"`
	CreatedBy bson.ObjectId `bson:"created_by,omitempty"`
	CreatedAt int64         `bson:"created_at"`
}

func (p *Post) Save() error {
	sess := db.CopySess()
	defer sess.Close()
	p.ID = bson.NewObjectId()
	err := sess.DB(conf.DBName).C(cPost).Insert(p)
	return err
}

func GetPost(query bson.M) (*Post, error) {
	sess := db.CopySess()
	defer sess.Close()
	post := &Post{}
	err := sess.DB(conf.DBName).C(cPost).Find(query).One(post)
	return post, err
}

func ListPosts(query bson.M) ([]*Post, error) {
	sess := db.CopySess()
	defer sess.Close()
	posts := []*Post{}
	err := sess.DB(conf.DBName).C(cPost).Find(query).Sort("-_id").All(posts)
	return posts, err
}

func ListPostsWithConfig(query bson.M, config *QueryConfig) ([]*Post, error) {
	sess := db.CopySess()
	defer sess.Close()
	posts := []*Post{}
	err := sess.DB(conf.DBName).C(cPost).Find(query).
		Skip(config.Skip).Limit(config.Limit).Sort("-_id").All(&posts)
	return posts, err
}

func CountPosts(query bson.M) (int, error) {
	sess := db.CopySess()
	defer sess.Close()
	return sess.DB(conf.DBName).C(cPost).Find(query).Count()
}
