package services

import (
	"database/sql"

	"github.com/insisthzr/blog-back/models"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u *User) ToModel() *models.User {
	return &models.User{
		ID:       u.ID,
		Name:     u.Name,
		Password: u.Password,
	}
}

func newUserFromModel(user *models.User) *User {
	return &User{
		ID:       user.ID,
		Name:     user.Name,
		Password: user.Password,
	}
}

func (u *User) Signup() error {
	tx, err := models.DB().Begin()
	if err != nil {
		return err
	}
	_, err = models.GetUserByName(tx, u.Name)
	if err == nil {
		return ErrUserExist
	}
	if err != sql.ErrNoRows {
		return err
	}
	user := u.ToModel()
	err = user.Insert(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	u.ID = user.ID
	user.Password = ""
	return nil
}

func (u *User) Login() error {
	tx, err := models.DB().Begin()
	user, err := models.GetUserByName(tx, u.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrUserNotExist
		}
		return err
	}
	tx.Commit()
	if u.Password != user.Password {
		return ErrNamePasswordError
	}
	u.ID = user.ID
	u.Password = ""
	return nil
}
