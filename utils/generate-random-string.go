package utils

import (
	"math/rand"
	"time"
)

const __random_str_charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func GenerateRandomString(length int) string {
	seededRnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	buffer := make([]byte, length)
	for i := range buffer {
		buffer[i] = __random_str_charset[seededRnd.Intn(len(__random_str_charset))]
	}

	return string(buffer)
}
