package helper

import (
	"math/rand"
	"strings"
)

func RandomString(name string, length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	res := string(b) + "_" + strings.ToLower(name)
	return res
}
