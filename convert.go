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
func CInt(s string) (ret int) {
	ret, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return
}

// CDate returns the result of time.ParseInLocation.
func CDate(s string) (ret time.Time) {
	ret, err := time.ParseInLocation("2006-1-2 15:04:05", s, time.Local)
	if err != nil {
		ret, _ = time.ParseInLocation("2006-01-02 15:04:05",
			"1900-01-01 00:00:00",
			time.Local)
	}
	return
}
