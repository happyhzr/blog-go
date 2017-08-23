package model

import (
	"github.com/insisthzr/blog-back/service/mysql"
	"github.com/insisthzr/blog-back/util"
)

type Post struct {
	Id        int64
	Title     string
	Content   string
	CreatedAt int64
	UpdatedAt int64
}

func (p *Post) Insert() {
	db := mysql.GetDb()
	stmt, err := db.Prepare("INSERT post SET title=?,content=?,created_at=?,updated_at=?")
	util.CheckError(err)
	result, err := stmt.Exec(p.Title, p.Content, p.CreatedAt, p.UpdatedAt)
	util.CheckError(err)
	id, err := result.LastInsertId()
	util.CheckError(err)
	p.Id = id
}

func ListPosts() []*Post {
	db := mysql.GetDb()
	stmt, err := db.Prepare("SELECT * FROM post ORDER BY updated_at DESC")
	util.CheckError(err)
	rows, err := stmt.Query()
	util.CheckError(err)
	posts := []*Post{}
	for rows.Next() {
		post := &Post{}
		err = rows.Scan(&post.Id, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)
		util.CheckError(err)
		posts = append(posts, post)
	}
	return posts
}
