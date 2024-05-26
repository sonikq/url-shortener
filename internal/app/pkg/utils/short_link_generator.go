package utils

import (
	"math/rand"
	"time"
)

// RandomString -
func RandomString(size int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	chars := []rune("ABCDEFGHIJKLMNOPRSTUVWXYZ" +
		"abcdefghijklmnoprstuvwxyz" +
		"0123456789")

	res := make([]rune, size)
	for i := range res {
		res[i] = chars[rnd.Intn(len(chars))]
	}

	return string(res)
}
