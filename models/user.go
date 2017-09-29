package models

import (
	"database/sql"
)

type User struct {
	ID       int64
	Name     string
	Password string
}

func (u *User) Insert(tx *sql.Tx) error {
	query := "INSERT INTO user(name, password) VALUES(?, ?)"
	res, err := tx.Exec(query, u.Name, u.Password)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = id
	return err
}

func getUser(tx *sql.Tx, query string, args ...interface{}) (*User, error) {
	u := &User{}
	err := tx.QueryRow(query, args...).Scan(&u.ID, &u.Name, &u.Password)
	return u, err
}

func GetUserByName(tx *sql.Tx, name string) (*User, error) {
	query := "SELECT id, name, password FROM user WHERE name = ?"
	user, err := getUser(tx, query, name)
	return user, err
}

func GetUserByID(tx *sql.Tx, id int) (*User, error) {
	query := "SELECT id, name, password FROM user WHERE id = ?"
	user, err := getUser(tx, query, id)
	return user, err
}
