package models

import "time"

type Post struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
	CreatedBy *User  `json:"created_by"`
}

func (p *Post) Insert() error {
	if p.CreatedAt == 0 {
		p.CreatedAt = time.Now().Unix()
	}
	db := getDb()
	query := "INSERT INTO posts(title, content, created_at, created_by) VALUES($1, $2, $3, $4) RETURNING id"
	err := db.QueryRow(query, p.Title, p.Content, p.CreatedAt, p.CreatedBy.ID).Scan(&p.ID)
	return err
}

func listPosts(query string, args ...interface{}) ([]*Post, error) {
	posts := []*Post{}
	db := getDb()
	rows, err := db.Query(query, args...)
	if err != nil {
		return posts, err
	}
	defer rows.Close()
	for rows.Next() {
		post := &Post{CreatedBy: &User{}}
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.CreatedBy.ID, &post.CreatedBy.Name)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	err = rows.Err()
	return posts, err
}

func ListPosts(offset int, limit int) ([]*Post, error) {
	query := `SELECT posts.id, title, content, created_at, users.id, users.name
	 FROM posts LEFT OUTER JOIN users ON posts.created_by = users.id 
	 ORDER BY posts.id DESC OFFSET $1`
	var posts []*Post
	var err error
	if limit == -1 {
		posts, err = listPosts(query, offset)
	} else {
		query += " LIMIT $2"
		posts, err = listPosts(query, offset, limit)
	}
	return posts, err
}

func countPost(query string, args ...interface{}) (int, error) {
	count := 0
	db := getDb()
	err := db.QueryRow(query, args...).Scan(&count)
	return count, err
}

func CountPost() (int, error) {
	query := "SELECT COUNT(*) FROM posts"
	count, err := countPost(query)
	return count, err
}
