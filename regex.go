package my

import (
	"regexp"
)

// Test provides some common regular expression detection features.
func Test(str, pattern string) (ret bool) {
	var pa string
	switch pattern {
	case "idcard":
		pa = `^\d{15}$)|(\d{17}(?:\d|x|X)$`
	case "english":
		pa = "^[A-Za-z]+$"
	case "chinese":
		pa = "^[\u4e00-\u9fa5]+$"
	case "username":
		pa = `^[a-z][a-z0-9]{4,19}$`
	case "email":
		pa = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
	case "int":
		pa = `^[-\+]?\d+$`
	case "number":
		pa = `^\d+$`
	case "double":
		pa = `^[-\+]?\d+(\.\d+)?$`
	case "price":
		pa = `^\d+(\.\d+)?$`
	case "zip":
		pa = `^[1-9]\d{5}$`
	case "qq":
		pa = `^[1-9]\d{4,9}$`
	case "phone":
		pa = `^((\(\d{2,3}\))|(\d{3}\-))?(\(0\d{2,3}\)|0\d{2,3}-)?[1-9]\d{6,7}(\-\d{1,4})?$`
	case "mobile":
		pa = `^(13[0-9]|14[5|7]|15[0-9]|18[0-9]|199)\d{8}$`
	case "url":
		pa = `^(http|https|ftp):\/\/[A-Za-z0-9]+\.[A-Za-z0-9]+[\/=\?%\-&_~\@[\\]\':+!]*([^<>\"])*$`
	case "domain":
		pa = `^[A-Za-z0-9\-]+\.([A-Za-z]{2,4}|[A-Za-z]{2,4}\.[A-Za-z]{2})$`
	case "ip":
		pa = `^\d+\.\d+\.\d+\.\d+$`
	case "password":
		return Test(str, "[A-Z]") && Test(str, "[a-z]") && Test(str, "[0-9]") && len(str) >= 8
	default:
		pa = pattern
	}
	reg := regexp.MustCompile(pa)

	return reg.MatchString(str)
}
