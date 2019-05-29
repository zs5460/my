package my

import (
	"strconv"
	"time"
	"unsafe"
)

// CStr returns the result of strconv.Itoa.
func CStr(i int) string {
	return strconv.Itoa(i)
}

// CInt returns the result of strconv.Atoi.
func CInt(s string) int {
	ret, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return ret
}

// CDate returns the result of time.ParseInLocation.
func CDate(s string) time.Time {
	ret, err := time.ParseInLocation("2006-1-2 15:04:05", s, time.Local)
	if err != nil {
		ret, err := time.ParseInLocation("2006-1-2", s, time.Local)
		if err != nil {
			ret, _ = time.ParseInLocation("2006-01-02", "1900-01-01", time.Local)
		}
		return ret
	}
	return ret
}

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// BytesToString Convert a byte slice to a string.
func BytesToString(b []byte) string {
	return bytes2str(b)
}

// StringToBytes Convert a string to a byte slice.
func StringToBytes(s string) []byte {
	return str2bytes(s)
}
