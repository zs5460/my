package my

import "unicode/utf8"

// IsEmpty ...
func IsEmpty(s string) bool {
	return len(s) == 0
}

// Left returns a specified number of characters from the left side of a string.
func Left(s string, n int) (ret string) {
	if n < 1 {
		return ""
	}
	runes := []rune(s)
	if n >= len(runes) {
		return s
	}

	ret = string(runes[:n])
	return
}

// Right returns a specified number of characters from the right side of a string.
func Right(s string, n int) (ret string) {
	runes := []rune(s)
	if n < 1 {
		return ""
	}
	if n >= len(runes) {
		return s
	}

	ret = string(runes[len(runes)-n:])
	return
}

// Mid returns a specified number of characters from a string.
func Mid(s string, start, length int) (ret string) {
	if start < 0 {
		start = 0
	}
	if start > len(s) {
		return ""
	}
	if length < 1 {
		return ""
	}

	runes := []rune(s)
	if start > len(runes) {
		return ""
	}
	if start+length > len(runes) {
		return string(runes[start:])
	}
	return string(runes[start : start+length])

}

// Len returns the number of rune in a string.
func Len(s string) int {
	if s == "" {
		return 0
	}
	return utf8.RuneCountInString(s)
}
