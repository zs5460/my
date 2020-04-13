package my_test

import (
	"testing"

	. "github.com/zs5460/my"
)

func TestTest(t *testing.T) {
	var tests = []struct {
		str     string
		pattern string
		want    bool
	}{
		{"430001001005001", "idcard", false},
		{"430001200010050012", "idcard", true},
		{"43000120001005001x", "idcard", true},
		{"43000120001005001X", "idcard", true},
		{"4300012000100500123456", "idcard", false},

		{"zhousong", "english", true},
		{"zs5460", "english", false},

		{"中华人民共和国", "chinese", true},
		{"公元2000年", "chinese", false},

		{"zs5460", "username", true},
		{"Zs5460", "username", false},

		{"zs5460@gmail.com", "email", true},
		{"zs5460@gmail", "email", false},
		{"zs5460.gmail.com", "email", false},
		{"zs5460@gmail.com.cn", "email", true},

		{"123456", "zip", true},
		{"012345", "zip", false},
		{"1234567", "zip", false},

		{"274619", "qq", true},
		{"0123", "qq", false},
		{"12345678900", "qq", false},

		{"731-12345678", "phone", true},
		{"0731-12345678", "phone", true},
		{"0731-12345678-01", "phone", true},
		{"0731-123456789", "phone", false},

		{"13073112345", "mobile", true},
		{"19973112345", "mobile", true},
		{"12073112345", "mobile", false},

		{"http://www.54600.net/index.htm", "url", true},
		{"https://www.google.com/?q=test", "url", true},
		{"ftp://54600.net/", "url", true},
		{"123456", "url", false},

		{"1.2.4.8", "ip", true},
		{"114.114.114.114", "ip", true},
		{"123.456", "ip", false},

		{"Zs123456", "password", true},
		{"zs5460", "password", false},
		{"Zs5460", "password", false},
		{"zs12345678", "password", false},

		{"5460", `^\d{4}$`, true},
		{"54600", `^\d{4}$`, false},
	}

	for _, test := range tests {
		if got := Test(test.str, test.pattern); got != test.want {
			t.Errorf("Test(%q,%q) = %v, want %v",
				test.str,
				test.pattern,
				got,
				test.want)
		}
	}
}
