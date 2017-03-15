package controllers

import (
	"crypto/sha1"
	"math/rand"
)

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

//CreateRandomPassword create a new random password
func CreateRandomPassword() string {
	b := make([]rune, 10)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

//ConvertToSHA1 convert the input string in crypted SHA1 string
func ConvertToSHA1(input string) []byte {
	h := sha1.New()
	h.Write([]byte(input))
	return h.Sum(nil)
}
