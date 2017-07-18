package util

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
)

func Hash(password string, salt string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(password+salt)))
}

func GenSalt() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", b)
}
