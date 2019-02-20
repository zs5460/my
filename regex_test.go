package my

import (
	"testing"
)

func TestTest(t *testing.T) {
	var tests = []struct {
		str     string
		pattern string
		want    bool
	}{
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

		{"1234", "number", true},
		{"zs5460", "number", false},

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
