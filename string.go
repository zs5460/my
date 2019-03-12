package my

import (
	"strings"
	"unicode/utf8"
)

// IsEmpty returns a Boolean value indicating whether a string is blank.
func IsEmpty(s string) bool {
	return len(s) == 0
}

// Left returns a specified number of characters from the left side of a string.
func Left(s string, n int) string {
	if n < 1 {
		return ""
	}
	runes := []rune(s)
	if n >= len(runes) {
		return s
	}

	return string(runes[:n])

}

// Right returns a specified number of characters from the right side of a string.
func Right(s string, n int) string {
	runes := []rune(s)
	if n < 1 {
		return ""
	}
	if n >= len(runes) {
		return s
	}

	return string(runes[len(runes)-n:])
}

// Mid returns a specified number of characters from a string.
func Mid(s string, start, length int) string {
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

// Space returns a string consisting of the specified number of spaces.
func Space(count int) string {
	return strings.Repeat(" ", count)
}
