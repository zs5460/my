package my

// IsEmpty ...
func IsEmpty(s string) bool {
	return len(s) == 0
}

// Left ...
func Left(s string, n int) (ret string) {
	if n < 1 {
		return ""
	}
	if n >= len(s) {
		return s
	}
	ret = s[:n]
	return
}

// Right ...
func Right(s string, n int) (ret string) {

	if n < 1 {
		return ""
	}
	if n >= len(s) {
		return s
	}
	ret = s[len(s)-n:]
	return
}
