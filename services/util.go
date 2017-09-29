package services

import (
	"errors"
)

var (
	ErrUserNotExist      = errors.New("user not exist")
	ErrUserExist         = errors.New("user exist")
	ErrNamePasswordError = errors.New("name or password error")
)
