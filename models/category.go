package models

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (c *Category) Create() error {
	query := `INSERT INTO category(name) VALUES(?)`
	res, err := DB().Exec(query, c.Name)
	if err != nil {
		return err
	}
	c.ID, _ = res.LastInsertId()
	return nil
}

func GetCategoryByID(id int64) (*Category, error) {
	query := `SELECT id, name FROM category WHERE id = ?`
	category := &Category{}
	err := DB().Get(category, query, id)
	return category, err
}

func hasCategoryByID(id int64) (bool, error) {
	_, err := GetCategoryByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func getCategoryTX(tx *sqlx.Tx, id int64) (*Category, error) {
	query := `SELECT id, name FROM category WHERE id = ?`
	category := &Category{}
	err := tx.Get(category, query, id)
	return category, err
}
