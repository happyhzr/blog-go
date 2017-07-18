package busniess

import (
	"errors"
	"time"

	"github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2"

	"github.com/insisthzr/blog-back/model"
	"github.com/insisthzr/blog-back/util"
)

var (
	errUserExist         = errors.New("user exist")
	errUserNotExist      = errors.New("user not exist")
	errIncorrectPassword = errors.New("incorrect password")
)

type UserIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserOut struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"createdAt"`
}

func newUserOut(user *model.User) *UserOut {
	out := &UserOut{
		ID:        user.ID.Hex(),
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
	return out
}

type SignupIn struct {
	UserIn
}

type SignupOut struct {
	UserOut
}

func Signup(in *SignupIn) (o *SignupOut, e error) {
	defer func() {
		logrus.WithFields(logrus.Fields{"in": in, "out": o, "error": e}).Info("sign up")
	}()
	_, err := model.GetUserByEmail(in.Email)
	if err != nil && err != mgo.ErrNotFound {
		panic(err)
	}
	if err == nil {
		return nil, errUserExist
	}

	salt := util.GenSalt()
	user := &model.User{
		Email:     in.Email,
		Salt:      salt,
		Password:  util.Hash(in.Password, salt),
		CreatedAt: time.Now().Unix(),
	}
	err = user.Save()
	if err != nil {
		panic(err)
	}
	out := &SignupOut{
		UserOut: *newUserOut(user),
	}
	return out, nil
}

type LoginIn struct {
	UserIn
}

type LoginOut struct {
	UserOut
}

func Login(in *LoginIn) (*LoginOut, error) {
	user, err := model.GetUserByEmail(in.Email)
	if err != nil {
		return nil, errUserNotExist
	}
	hash := util.Hash(in.Password, user.Salt)
	if hash != user.Password {
		return nil, errIncorrectPassword
	}
	out := &LoginOut{
		UserOut: *newUserOut(user),
	}
	return out, nil
}
