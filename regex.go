package my

import (
	"regexp"
)

var commonPatterns = map[string]*regexp.Regexp{
	"idcard":   regexp.MustCompile(`^\d{17}(\d|x|X)$`),
	"english":  regexp.MustCompile(`^[A-Za-z]+$`),
	"chinese":  regexp.MustCompile("^[\u4e00-\u9fa5]+$"),
	"username": regexp.MustCompile(`^[a-z][a-z0-9]{4,19}$`),
	"email":    regexp.MustCompile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`),
	"zip":      regexp.MustCompile(`^[1-9]\d{5}$`),
	"qq":       regexp.MustCompile(`^[1-9]\d{4,9}$`),
	"phone":    regexp.MustCompile(`^((\(\d{2,3}\))|(\d{3}\-))?(\(0\d{2,3}\)|0\d{2,3}-)?[1-9]\d{6,7}(\-\d{1,4})?$`),
	"mobile":   regexp.MustCompile(`^(13[0-9]|14[57]|15[0-9]|18[0-9]|199)\d{8}$`),
	"url":      regexp.MustCompile(`^((ht|f)tps?):\/\/[\w\-]+(\.[\w\-]+)+([\w\-.,@?^=%&:\/~+#]*[\w\-@?^=%&\/~+#])?$`),
	"ip":       regexp.MustCompile(`^\d+\.\d+\.\d+\.\d+$`),
}

// Test provides some common regular expression detection features.
func Test(str, pattern string) bool {
	if pattern == "password" {
		return isStrongPassword(str)
	}

	reg, ok := commonPatterns[pattern]
	if !ok {
		var err error
		reg, err = regexp.Compile(pattern)
		if err != nil {
			return false
		}
	}
	return reg.MatchString(str)
}

func isStrongPassword(s string) bool {
	return len(s) >= 8 && Test(s, "[A-Z]") && Test(s, "[a-z]") && Test(s, "[0-9]")
}
