package models

import (
	"database/sql"
)

type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (t *Tag) Insert(tx *sql.Tx) error {
	query := "INSERT INTO tag(name) VALUES(?)"
	res, err := tx.Exec(query, t.Name)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = id
	return err
}

func listTags(tx *sql.Tx, query string, args ...interface{}) ([]*Tag, error) {
	tags := []*Tag{}
	rows, err := tx.Query(query, args...)
	if err != nil {
		return tags, err
	}
	defer rows.Close()

	for rows.Next() {
		tag := &Tag{}
		err := rows.Scan(&tag.ID, &tag.Name)
		if err != nil {
			break
		}
		tags = append(tags, tag)
	}
	err = rows.Err()
	if err != nil {
		return tags, err
	}
	return tags, nil
}

func ListTags(tx *sql.Tx, limit int, offset int) ([]*Tag, error) {
	query := "SELECT id, name from tag LIMIT ? OFFSET ?"
	tags, err := listTags(tx, query, limit, offset)
	return tags, err
}
