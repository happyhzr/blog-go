package services

import (
	"github.com/insisthzr/blog-back/models"
)

type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (t *Tag) ToModel() *models.Tag {
	return &models.Tag{
		ID:   t.ID,
		Name: t.Name,
	}
}

func newTagFromModel(tag *models.Tag) *Tag {
	return &Tag{
		ID:   tag.ID,
		Name: tag.Name,
	}
}

func (t *Tag) Save() error {
	tx, err := models.DB().Begin()
	if err != nil {
		return err
	}
	tag := t.ToModel()
	err = tag.Insert(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	t.ID = tag.ID
	return nil
}

func ListTags(limit int, offset int) ([]*Tag, error) {
	tags := []*Tag{}
	tx, err := models.DB().Begin()
	if err != nil {
		return tags, err
	}
	ts, err := models.ListTags(tx, limit, offset)
	if err != nil {
		tx.Rollback()
		return tags, err
	}
	tx.Commit()
	tags = make([]*Tag, 0, len(ts))
	for _, t := range ts {
		tags = append(tags, newTagFromModel(t))
	}
	return tags, nil
}
