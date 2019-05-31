package my

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RndNumber returns a string of numbers of the specified length.
func RndNumber(n int) string {
	return genRandom(n, 0)
}

// RndAlpha returns a string of letters of the specified length.
func RndAlpha(n int) string {
	return genRandom(n, 1)
}

// RndString returns a random string of the specified length.
func RndString(n int) string {
	return genRandom(n, 2)
}

// RndFilename returns a file name consisting of a time and a random string.
func RndFilename(ext string) string {
	if Left(ext, 1) != "." {
		ext = "." + ext
	}
	return time.Now().Format("20060102150405") + genRandom(6, 0) + ext
}

func genRandom(n, charType int) string {
	if n < 0 || n > 32 {
		n = 32
	}
	if charType < 0 || charType > 2 {
		charType = 1
	}
	chars := []string{
		"0123456789",
		"abcdefghijklmnopqrstuvwxyz",
		"abcdefghijklmnopqrstuvwxyz0123456789",
	}
	length := len(chars[charType])
	bytes := []byte(chars[charType])
	result := []byte{}

	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(length)])
	}
	return string(result)
}

// GenRandom ...
func GenRandom(n, charType int) string {
	return genRandom(n, charType)
}
