package my

import (
	"strconv"
	"time"
)

// CStr converts an integer to a string.
// It is a wrapper around strconv.Itoa.
func CStr(i int) string {
	return strconv.Itoa(i)
}

// CInt converts a string to an integer.
// It is a wrapper around strconv.Atoi.
// Returns 0 if the string cannot be converted to an integer.
func CInt(s string) int {
	ret, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return ret
}

// CDate parses a string into a time.Time object.
// It accepts formats "2006-1-2 15:04:05" and "2006-1-2".
// Returns 1900-01-01 if the string cannot be parsed.
func CDate(s string) time.Time {
	ret, err := time.ParseInLocation("2006-1-2 15:04:05", s, time.Local)
	if err != nil {
		ret, err = time.ParseInLocation("2006-1-2", s, time.Local)
		if err != nil {
			ret, _ = time.ParseInLocation("2006-01-02", "1900-01-01", time.Local)
		}
		return ret
	}
	return ret
}

// BytesToString converts a byte slice to a string.
// Deprecated: This function will be removed in a future version.
// Use the standard library's string() conversion directly instead.
func BytesToString(b []byte) string {
	return string(b)
}

// StringToBytes converts a string to a byte slice.
// Deprecated: This function will be removed in a future version.
// Use the standard library's []byte() conversion directly instead.
func StringToBytes(s string) []byte {
	return []byte(s)
}
