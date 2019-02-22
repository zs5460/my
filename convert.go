package my

import (
	"strconv"
	"time"
)

// CStr returns the result of strconv.Itoa.
func CStr(i int) string {
	return strconv.Itoa(i)
}

// CInt returns the result of strconv.Atoi.
func CInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// CDate returns the result of time.ParseInLocation.
func CDate(s string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
}
