package services

import (
	"github.com/insisthzr/blog-back/models"
)

type Category struct {
	ID   int64
	Name string
}

func (c *Category) ToModel() *models.Category {
	return &models.Category{
		ID:   c.ID,
		Name: c.Name,
	}
}

func newCategoryFromModel(category *models.Category) *Category {
	return &Category{
		ID:   category.ID,
		Name: category.Name,
	}
}

func (c *Category) Save() error {
	tx, err := models.DB().Begin()
	if err != nil {
		return err
	}
	category := c.ToModel()
	err = category.Insert(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	c.ID = category.ID
	return nil
}

func ListCategorys(limit int, offset int) ([]*Category, error) {
	categorys := []*Category{}
	tx, err := models.DB().Begin()
	if err != nil {
		return categorys, err
	}
	cs, err := models.ListCategorys(tx, limit, offset)
	if err != nil {
		tx.Rollback()
		return categorys, err
	}
	tx.Commit()
	categorys = make([]*Category, 0, len(cs))
	for _, c := range cs {
		categorys = append(categorys, newCategoryFromModel(c))
	}
	return categorys, nil
}
