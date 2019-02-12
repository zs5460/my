package my

// Left ...
func Left(s string, n int) (ret string) {
	if len(s) < n {
		ret = s
	}
	ret = s[:n]
	return
}

// Right ...
func Right(s string, n int) (ret string) {

	if n < 1 || len(s) < n {
		ret = s
		return
	}
	ret = s[len(s)-n:]
	return
}
