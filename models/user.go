package models

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID       int64  `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Password string `db:"password" json:"password"`
}

func (u *User) Signup() error {
	tx := DB().MustBegin()
	user := User{}
	query := `SELECT id, name, password FROM user WHERE name = ? LIMIT 1`
	err := tx.Get(&user, query, u.Name)
	if err == nil {
		tx.Rollback()
		return errors.New("user exist")
	}
	query = `INSERT INTO user(name, password) VALUES(?, ?)`
	res, err := tx.Exec(query, u.Name, u.Password)
	if err != nil {
		tx.Rollback()
		return err
	}
	u.ID, err = res.LastInsertId()
	tx.Commit()
	return err
}

func (u *User) Login() error {
	user := User{}
	query := `SELECT id, name, password FROM user WHERE name = ? LIMIT 1`
	err := DB().Get(&user, query, u.Name)
	if err != nil {
		return errors.New("username not exist")
	}
	u.ID = user.ID
	if u.Password != user.Password {
		return errors.New("username or password incorrect")
	}
	return nil
}

func getUserByID(id int64) (*User, error) {
	user := &User{}
	query := `SELECT id, name FROM user WHERE id = ? LIMIT 1`
	err := DB().Get(user, query, id)
	return user, err
}

func listUserByIDs(ids []int64) ([]*User, error) {
	users := []*User{}
	query := `SELECT id, name FROM user WHERE id IN (?)`
	query, args, err := sqlx.In(query, ids)
	err = DB().Select(&users, query, args...)
	return users, err
}

func getUserMap(ids []int64) (map[int64]*User, error) {
	users, err := listUserByIDs(ids)
	if err != nil {
		return nil, err
	}
	m := map[int64]*User{}
	for _, u := range users {
		m[u.ID] = u
	}
	return m, nil
}
