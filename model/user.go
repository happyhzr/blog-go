package model

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/insisthzr/blog-back/conf"
	"github.com/insisthzr/blog-back/db"
)

type User struct {
	ID        bson.ObjectId `bson:"_id"`
	Email     string        `bson:"email"`
	Password  string        `bson:"password"`
	Salt      string        `bson:"salt"`
	CreatedAt int64         `bson:"created_at"`
}

func (u *User) Save() error {
	sess := db.CopySess()
	defer sess.Close()
	u.ID = bson.NewObjectId()
	err := sess.DB(conf.DBName).C(cUser).Insert(u)
	return err
}

func GetUser(query bson.M) (*User, error) {
	sess := db.CopySess()
	defer sess.Close()
	user := &User{}
	err := sess.DB(conf.DBName).C(cUser).Find(query).One(user)
	return user, err
}

func GetUserByEmail(email string) (*User, error) {
	query := bson.M{"email": email}
	return GetUser(query)
}

func GetUserByID(id bson.ObjectId) (*User, error) {
	query := bson.M{"_id": id}
	return GetUser(query)
}
