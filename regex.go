package my

import (
	"regexp"
)

// Test provides some common regular expression detection features.
func Test(str, pattern string) bool {
	var pa string
	switch pattern {
	case "idcard":
		pa = `(^\d{15}$)|(^\d{17}(\d|x|X)$)`
	case "english":
		pa = "^[A-Za-z]+$"
	case "chinese":
		pa = "^[\u4e00-\u9fa5]+$"
	case "username":
		pa = `^[a-z][a-z0-9]{4,19}$`
	case "email":
		pa = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
	case "zip":
		pa = `^[1-9]\d{5}$`
	case "qq":
		pa = `^[1-9]\d{4,9}$`
	case "phone":
		pa = `^((\(\d{2,3}\))|(\d{3}\-))?(\(0\d{2,3}\)|0\d{2,3}-)?[1-9]\d{6,7}(\-\d{1,4})?$`
	case "mobile":
		pa = `^(13[0-9]|14[5|7]|15[0-9]|18[0-9]|199)\d{8}$`
	case "url":
		pa = `^((ht|f)tps?):\/\/[\w\-]+(\.[\w\-]+)+([\w\-.,@?^=%&:\/~+#]*[\w\-@?^=%&\/~+#])?$`
	case "ip":
		pa = `^\d+\.\d+\.\d+\.\d+$`
	case "password":
		return isStrongPassword(str)
	default:
		pa = pattern
	}
	reg := regexp.MustCompile(pa)

	return reg.MatchString(str)
}

func isStrongPassword(s string) bool {
	return Test(s, "[A-Z]") && Test(s, "[a-z]") && Test(s, "[0-9]") && len(s) >= 8
}
