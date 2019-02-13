package my

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RndStr ...
func RndStr(n int) string {
	if n < 1 {
		return ""
	}
	if n > 32 {
		n = 32
	}
	str := "abcdefghijklmnopqrstuvwxyz0123456789"
	length := len(str)
	bytes := []byte(str)
	result := []byte{}

	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(length)])
	}
	return string(result)
}
