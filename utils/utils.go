package utils

import (
	"math/rand"
)

func GenerateToken(n int) string {
	var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var number = []byte("0123456789")
	var alphaNumeric = append(charset, number...)

	b := make([]byte, n)
	for i := range b {
		// randomly select 1 character from given charset
		b[i] = alphaNumeric[rand.Intn(len(alphaNumeric))]
	}
	return string(b)
}
