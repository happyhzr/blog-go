package models

import (
	"github.com/jmoiron/sqlx"

	"github.com/insisthzr/blog-back/utils"
)

type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (t *Tag) Create() error {
	query := `INSERT INTO tag(name) VALUES(?)`
	res, err := DB().Exec(query, t.Name)
	if err != nil {
		return err
	}
	t.ID, _ = res.LastInsertId()
	return nil
}

func listTagByIDs(ids []int64) ([]*Tag, error) {
	tags := []*Tag{}
	query := `SELECT id, name FROM tag WHERE id IN (?)`
	query, args, err := sqlx.In(query, ids)
	if err != nil {
		return []*Tag{}, err
	}
	query = DB().Rebind(query)
	err = DB().Select(&tags, query, args...)
	return tags, err
}

func listTagByPostID(id int64) ([]*Tag, error) {
	pts, err := listPostTagByPostID(id)
	if err != nil {
		return []*Tag{}, err
	}
	ids := make([]int64, 0, len(pts))
	for _, pt := range pts {
		ids = append(ids, pt.TagID)
	}
	tags, err := listTagByIDs(ids)
	return tags, err
}

func getTagMapByPostID(id int64) (map[int64]*Tag, error) {
	tags, err := listTagByPostID(id)
	if err != nil {
		return map[int64]*Tag{}, err
	}
	m := map[int64]*Tag{}
	for _, t := range tags {
		m[t.ID] = t
	}
	return m, nil
}

func hasTagByIDs(ids []int64) (bool, error) {
	tags, err := listTagByIDs(ids)
	if err != nil {
		return false, err
	}
	tagIDs := make([]int64, 0, len(tags))
	for _, t := range tags {
		tagIDs = append(tagIDs, t.ID)
	}
	return utils.IsEqualInt64s(ids, tagIDs), nil
}
