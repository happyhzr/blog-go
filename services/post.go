package services

import (
	"time"

	"github.com/insisthzr/blog-back/models"
	"github.com/insisthzr/blog-back/utils"
)

type Post struct {
	ID         int64    `json:"id"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	CreatedAt  int64    `json:"createdAt"`
	AuthorID   int64    `json:"authorID"`
	Author     User     `json:"author"`
	TagIDs     []int64  `json:"tagIDs"`
	Tags       []Tag    `json:"tags"`
	CategoryID int64    `json:"categoryID"`
	Category   Category `json:"category"`
}

func (p *Post) ToModel() *models.Post {
	return &models.Post{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: time.Unix(p.CreatedAt, 0),
		Author: models.User{
			ID: p.AuthorID,
		},
		Category: models.Category{
			ID: p.CategoryID,
		},
	}
}

func newPostFromModel(post *models.Post) *Post {
	return &Post{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt.Unix(),
		Author: User{
			ID:   post.Author.ID,
			Name: post.Author.Name,
		},
		Category: Category{
			ID:   post.Category.ID,
			Name: post.Category.Name,
		},
	}
}

func (p *Post) Save() error {
	tx, err := models.DB().Begin()
	if err != nil {
		return err
	}
	post := p.ToModel()
	err = post.Insert(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	p.ID = post.ID
	pts := make([]*models.PostTag, 0, len(p.TagIDs))
	for _, id := range p.TagIDs {
		pts = append(pts, &models.PostTag{PostID: p.ID, TagID: id})
	}
	err = models.PostTags(pts).Insert(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (p *Post) Update(id int64) error {
	tx, err := models.DB().Begin()
	if err != nil {
		return err
	}
	post := p.ToModel()
	err = post.Update(tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func listPosts(s models.PostSelector) ([]*Post, error) {
	posts := []*Post{}
	tx, err := models.DB().Begin()
	if err != nil {
		return posts, err
	}
	ps, err := models.ListPosts(tx, s)
	if err != nil {
		tx.Rollback()
		return posts, err
	}
	tx.Commit()
	for _, p := range ps {
		posts = append(posts, newPostFromModel(p))
	}
	return posts, nil
}

func ListPostsS(s models.PostSelector) ([]*Post, error) {
	return listPosts(s)
}

func ListPosts(limit int, offset int) ([]*Post, error) {
	s := models.PostSelector{
		LimitOffset: utils.LimitOffset{
			Limit:  limit,
			Offset: offset,
		},
	}
	return listPosts(s)
}

func ListPostsByCategory(cID int64, limit int, offset int) ([]*Post, error) {
	s := models.PostSelector{
		CID: cID,
		LimitOffset: utils.LimitOffset{
			Limit:  limit,
			Offset: offset,
		},
	}
	return listPosts(s)
}

func GetPost(id int64) (*Post, error) {
	tx, err := models.DB().Begin()
	if err != nil {
		return nil, err
	}
	p, err := models.GetPost(tx, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return newPostFromModel(p), nil
}

func DeletePost(id int64) error {
	tx, err := models.DB().Begin()
	if err != nil {
		return err
	}
	err = models.DeletePost(tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
