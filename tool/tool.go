package tool

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func IsEqualInt64s(a []int64, b []int64) bool {
	for _, i := range a {
		found := false
		for _, j := range b {
			if i == j {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func MakeSalt() string {
	str := "asdfghjkl0123456789"
	n := len(str)
	res := ""
	for i := 0; i < 16; i++ {
		res += string(str[rand.Intn(n)])
	}
	return res
}

func HashPassword(password string, salt string) string {
	t := md5.New()
	io.WriteString(t, password+salt)
	return fmt.Sprintf("%x\n", t.Sum(nil))
}
