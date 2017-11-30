package models

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type PostTag struct {
	ID     int64 `db:"id"`
	PostID int64 `db:"post_id"`
	TagID  int64 `db:"tag_id"`
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

func listPostTagByPostID(id int64) ([]*PostTag, error) {
	pts := []*PostTag{}
	query := `SELECT id, post_id, tag_id FROM post_tag WHERE post_id = ?`
	err := DB().Select(&pts, query, id)
	return pts, err
}

type PostTags []*PostTag

func (pts PostTags) createTX(tx *sqlx.Tx) error {
	query := "INSERT INTO post_tag(post_id, tag_id) VALUES "
	vals := []interface{}{}
	for i, pt := range pts {
		if i != 0 {
			query += ","
		}
		query += "(?, ?)"
		vals = append(vals, pt.PostID, pt.TagID)
	}
	_, err := tx.Exec(query, vals...)
	return err
}
