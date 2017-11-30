package models

import (
	"database/sql"
	"errors"
	"time"
)

type Post struct {
	ID         int64     `db:"id" json:"id"`
	Title      string    `db:"title" json:"title"`
	Content    string    `db:"content" json:"content"`
	CreatedAt  int64     `db:"created_at" json:"created_at"`
	UpdateAt   int64     `db:"updated_at" json:"updated_at"`
	AuthorID   int64     `db:"author_id" json:"-"`
	CategoryID int64     `db:"category_id" json:"-"`
	Author     *User     `json:"author"`
	Category   *Category `json:"category"`
	Tags       []*Tag    `json:"tags"`
}

func (p *Post) Create() error {
	err := p.checkCreate()
	if err != nil {
		return err
	}
	return p.create()
}

func (p *Post) checkCreate() error {
	var categoryID int64
	if p.Category != nil {
		categoryID = p.Category.ID
	}
	has, err := hasCategoryByID(categoryID)
	if err != nil {
		return err
	}
	if !has {
		return errors.New("category id not exist")
	}

	tagIDs := make([]int64, 0, len(p.Tags))
	for _, t := range p.Tags {
		tagIDs = append(tagIDs, t.ID)
	}
	has, err = hasTagByIDs(tagIDs)
	if err != nil {
		return err
	}
	if !has {
		return errors.New("tag ids not exist")
	}
	return nil
}

func (p *Post) create() error {
	var aID int64
	if p.Author != nil {
		aID = p.Author.ID
	}
	var cID int64
	if p.Category != nil {
		cID = p.Category.ID
	}
	t := time.Now().Unix()
	p.CreatedAt = t
	p.UpdateAt = t

	query := `INSERT INTO post(title, content, author_id, category_id, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)`
	tx, err := DB().Beginx()
	if err != nil {
		return err
	}
	res, err := tx.Exec(query, p.Title, p.Content, aID, cID, p.CreatedAt, p.UpdateAt)
	if err != nil {
		tx.Rollback()
		return err
	}
	p.ID, _ = res.LastInsertId()
	if len(p.Tags) != 0 {
		pts := make(PostTags, 0, len(p.Tags))
		for _, t := range p.Tags {
			pts = append(pts, &PostTag{PostID: p.ID, TagID: t.ID})
		}
		err = pts.createTX(tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func (p *Post) Update() error {
	err := p.checkUpdate()
	if err != nil {
		return err
	}
	return p.update()
}

func (p *Post) checkUpdate() error {
	post, err := getPostByID(p.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("post not exist")
		}
		return err
	}
	if p.AuthorID != post.AuthorID {
		return errors.New("author id not equal")
	}
	return nil
}

func (p *Post) update() error {
	query := `UPDATE post SET title = ?, content = ?, updated_at = ? WHERE id = ?`
	t := time.Now().Unix()
	_, err := DB().Exec(query, p.Title, p.Content, t, p.ID)
	return err
}

func (p *Post) Delete() error {
	err := p.checkDelete()
	if err != nil {
		return err
	}
	return p.delete()
}

func (p *Post) checkDelete() error {
	post, err := getPostByID(p.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("post not exist")
		}
		return err
	}
	if p.AuthorID != post.AuthorID {
		return errors.New("author id not equal")
	}
	return nil
}

func (p *Post) delete() error {
	query := `DELETE FROM post WHERE id = ?`
	_, err := DB().Exec(query, p.ID)
	return err
}

func (p *Post) fillAssociation() error {
	user, err := getUserByID(p.AuthorID)
	if err == nil {
		p.Author = user
	} else if err != sql.ErrNoRows {
		return err
	}

	category, err := GetCategoryByID(p.CategoryID)
	if err == nil {
		p.Category = category
	} else if err != sql.ErrNoRows {
		return err
	}

	tags, err := listTagByPostID(p.ID)
	if err == nil {
		p.Tags = tags
	} else if err != sql.ErrNoRows {
		return err
	}

	return nil
}

func ListPostWithRange(offset int, limit int) ([]*Post, error) {
	posts := []*Post{}
	query := `SELECT id, title, content, created_at, updated_at, author_id, category_id FROM post
	 ORDER BY updated_at DESC LIMIT ? OFFSET ?`
	err := DB().Select(&posts, query, limit, offset)
	for _, p := range posts {
		err = p.fillAssociation()
		if err != nil {
			return nil, err
		}
	}
	return posts, err
}

func GetPostByID(id int64) (*Post, error) {
	post, err := getPostByID(id)
	if err != nil {
		return nil, err
	}
	err = post.fillAssociation()
	return post, err
}

func getPostByID(id int64) (*Post, error) {
	post := &Post{}
	query := `SELECT id, title, content, created_at, updated_at, author_id, category_id FROM post WHERE id = ? LIMIT 1`
	err := DB().Get(post, query, id)
	return post, err
}
