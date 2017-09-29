package models

import (
	"database/sql"
)

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (c *Category) Insert(tx *sql.Tx) error {
	query := "INSERT INTO category(name) VALUES(?)"
	res, err := tx.Exec(query, c.Name)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	c.ID = id
	return err
}

func listCategorys(tx *sql.Tx, query string, args ...interface{}) ([]*Category, error) {
	categorys := []*Category{}
	rows, err := tx.Query(query, args...)
	if err != nil {
		return categorys, err
	}
	defer rows.Close()

	for rows.Next() {
		c := &Category{}
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			break
		}
		categorys = append(categorys, c)
	}
	if rows.Err() != nil {
		return categorys, err
	}
	return categorys, nil
}

func ListCategorys(tx *sql.Tx, limit int, offset int) ([]*Category, error) {
	query := "SELECT id, name from category LIMIT ? OFFSET ?"
	categorys, err := listCategorys(tx, query, limit, offset)
	return categorys, err
}
