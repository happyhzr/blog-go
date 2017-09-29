package models

import (
	"database/sql"
)

type PostTag struct {
	ID     int64
	PostID int64
	TagID  int64
}

func (p *PostTag) Insert(tx sql.Tx) error {
	query := "INSERT INTO post_tag(post_id, tag_id) VALUES(?, ?)"
	res, err := tx.Exec(query, p.PostID, p.TagID)
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

type PostTags []*PostTag

func (pts PostTags) Insert(tx *sql.Tx) error {
	query := "INSERT INTO post_tag(post_id, tag_id) VALUES "
	vals := []interface{}{}
	for _, pt := range pts {
		query += "(?, ?),"
		vals = append(vals, pt.PostID, pt.TagID)
	}
	query = query[:len(query)-1]
	_, err := tx.Exec(query, vals...)
	return err
}
