package util

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
)

var (
	SaltLen = 16
)

func Hash(password string, salt string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(password+salt)))
}

func GenSalt() string {
	b := make([]byte, SaltLen)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", b)
}
