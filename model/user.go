package model

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type User struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) AfterFind() (err error) {
	u.CreatedAtTS = u.CreatedAt.Unix()
	u.UpdatedAtTS = u.UpdatedAt.Unix()
	return nil
}

func (u *User) Signup() error {
	exist := &User{}
	tx := GetDB().Begin()
	err := tx.Where("username = ?", u.Username).First(exist).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			tx.Rollback()
			return err
		}
		return tx.Commit().Error
	}
	return errors.New("user exit")
}

func (u *User) Login() error {
	exist := &User{}
	err := GetDB().Where("username = ?", u.Username).First(exist).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("user not exist")
		}
		return err
	}
	if u.Password != exist.Password {
		return errors.New("username or password error")
	}
	u.CreatedAtTS = exist.CreatedAtTS
	u.UpdatedAtTS = exist.CreatedAtTS
	return nil
}
