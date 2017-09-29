package models

import (
	"time"

	"database/sql"

	"github.com/insisthzr/blog-back/utils"
)

type PostSelector struct {
	utils.LimitOffset
	CID int64 `form:"cid"`
}

type Post struct {
	ID        int64
	Title     string
	Content   string
	CreatedAt time.Time
	Author    User
	Category  Category
}

func (p *Post) Insert(tx *sql.Tx) error {
	query := "INSERT INTO post(title, content, created_at, author_id, category_id) VALUES(?, ?, ?, ?, ?)"
	res, err := tx.Exec(query, p.Title, p.Content, p.CreatedAt, p.Author.ID, p.Category.ID)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = id
	return nil
}

func (p *Post) Update(tx *sql.Tx, id int64) error {
	query := "UPDATE post SET title = ?, content = ? WHERE id = ?"
	_, err := tx.Exec(query, p.Title, p.Content, id)
	return err
}

func getPost(tx *sql.Tx, query string, args ...interface{}) (*Post, error) {
	p := &Post{}
	authorName := sql.NullString{}
	categoryName := sql.NullString{}
	err := tx.QueryRow(query, args...).Scan(&p.ID, &p.Title, &p.Content, &p.CreatedAt, &p.Author.ID, &authorName, &p.Category.ID, &categoryName)
	if err != nil {
		return nil, err
	}
	p.Author.Name = authorName.String
	p.Category.Name = categoryName.String
	return p, err
}

func GetPost(tx *sql.Tx, id int64) (*Post, error) {
	query := `SELECT post.id, title, content, created_at, author_id, user.name, category_id, category.name 
	 FROM post
	 LEFT OUTER JOIN user ON author_id = user.id 
	 LEFT OUTER JOIN category ON category_id = category.id
     WHERE post.id = ?`
	return getPost(tx, query, id)
}

func listPosts(tx *sql.Tx, query string, args ...interface{}) ([]*Post, error) {
	posts := []*Post{}
	rows, err := tx.Query(query, args...)
	if err != nil {
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		p := &Post{}
		authorName := sql.NullString{}
		categoryName := sql.NullString{}
		err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.CreatedAt, &p.Author.ID, &authorName, &p.Category.ID, &categoryName)
		if err != nil {
			break
		}
		p.Author.Name = authorName.String
		p.Category.Name = categoryName.String
		posts = append(posts, p)
	}
	err = rows.Err()
	if err != nil {
		return posts, err
	}
	return posts, err
}

func ListPosts(tx *sql.Tx, s PostSelector) ([]*Post, error) {
	ps := []*Post{}
	var err error
	if s.CID == 0 {
		query := `SELECT post.id, title, content, created_at, author_id, user.name, category_id, category.name FROM post
		 LEFT OUTER JOIN user ON author_id = user.id 
		 LEFT OUTER JOIN category ON category_id = category.id
		 ORDER BY post.id DESC LIMIT ? OFFSET ?`
		ps, err = listPosts(tx, query, s.Limit, s.Offset)
	} else {
		query := `SELECT post.id, title, content, created_at, author_id, user.name, category_id, category.name FROM post
		 LEFT OUTER JOIN user ON author_id = user.id 
		 LEFT OUTER JOIN category ON category_id = category.id
		 WHERE post.category_id = ?
		 ORDER BY post.id DESC LIMIT ? OFFSET ?`
		ps, err = listPosts(tx, query, s.CID, s.Limit, s.Offset)
	}
	return ps, err
}

func DeletePost(tx *sql.Tx, id int64) error {
	query := "DELETE from post WHERE id = ?"
	_, err := tx.Exec(query, id)
	return err
}
