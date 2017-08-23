package model

import (
	"github.com/insisthzr/blog-back/service/mysql"
	"github.com/insisthzr/blog-back/util"
)

type User struct {
	Id       int64
	Username string
	Password string
}

func (u *User) Insert() {
	db := mysql.GetDb()
	stmt, err := db.Prepare("INSERT user SET username=?,password=?")
	util.CheckError(err)
	result, err := stmt.Exec(u.Username, u.Password)
	util.CheckError(err)
	id, err := result.LastInsertId()
	util.CheckError(err)
	u.Id = id
}

func GetUserByUsername(username string) *User {
	db := mysql.GetDb()
	stmt, err := db.Prepare("SELECT * FROM user WHERE username=? LIMIT 1")
	util.CheckError(err)
	rows, err := stmt.Query(username)
	util.CheckError(err)
	ok := rows.Next()
	if !ok {
		return nil
	}
	user := &User{}
	err = rows.Scan(&user.Id, &user.Username, &user.Password)
	util.CheckError(err)
	return user
}
