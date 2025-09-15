package my

import (
	"strings"
	"unicode/utf8"
)

// IsEmpty checks if a string is empty.
// Returns true if the string length is 0, false otherwise.
func IsEmpty(s string) bool {
	return len(s) == 0
}

// Left returns the first n characters from the left side of a string.
// If n is less than 1, returns an empty string.
// If n is greater than or equal to the string length, returns the entire string.
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

// Right returns the last n characters from the right side of a string.
// If n is less than 1, returns an empty string.
// If n is greater than or equal to the string length, returns the entire string.
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

// Mid returns a substring starting at the specified position with the specified length.
// If start is negative, it is treated as 0.
// If start is greater than the string length, returns an empty string.
// If length is less than 1, returns an empty string.
// If start+length exceeds the string length, returns the substring from start to the end.
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

// Len returns the number of Unicode code points (runes) in a string.
// This is different from len(s) which returns the number of bytes.
func Len(s string) int {
	if s == "" {
		return 0
	}
	return utf8.RuneCountInString(s)
}

// Space returns a string consisting of the specified number of spaces.
// It is a wrapper around strings.Repeat(" ", count).
func Space(count int) string {
	return strings.Repeat(" ", count)
}
