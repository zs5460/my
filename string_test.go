package my

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	var tests = []struct {
		src  string
		want bool
	}{
		{"zs5460", false},
		{"", true},
	}

	for _, test := range tests {
		if got := IsEmpty(test.src); got != test.want {
			t.Errorf("IsEmpty(%q) = %v, want %v",
				test.src,
				got,
				test.want)
		}
	}
}

func TestLeft(t *testing.T) {
	var tests = []struct {
		src    string
		length int
		want   string
	}{
		{"zs5460", 2, "zs"},
		{"zs5460", 8, "zs5460"},
		{"zs5460", 1, "z"},
		{"zs5460", 0, ""},
		{"zs5460", -1, ""},
		{"", 2, ""},
		{"中华人民共和国", 2, "中华"},
		{"公元2000年", 5, "公元200"},
	}

	for _, test := range tests {
		if got := Left(test.src, test.length); got != test.want {
			t.Errorf("Left(%q,%d) = %v, want %s",
				test.src,
				test.length,
				got,
				test.want)
		}
	}
}

func TestRight(t *testing.T) {

	var tests = []struct {
		src    string
		length int
		want   string
	}{
		{"zs5460", 2, "60"},
		{"zs5460", 8, "zs5460"},
		{"zs5460", 1, "0"},
		{"zs5460", 0, ""},
		{"zs5460", -1, ""},
		{"", 2, ""},
		{"中华人民共和国", 3, "共和国"},
		{"公元2000年", 5, "2000年"},
	}

	for _, test := range tests {
		if got := Right(test.src, test.length); got != test.want {
			t.Errorf("Right(%q,%d) = %v, want %s",
				test.src,
				test.length,
				got,
				test.want)
		}
	}

}

func TestMid(t *testing.T) {

	var tests = []struct {
		src    string
		start  int
		length int
		want   string
	}{
		{"zs5460", -1, 2, "zs"},
		{"zs5460", 0, 2, "zs"},
		{"zs5460", 2, 2, "54"},
		{"zs5460", 4, 2, "60"},
		{"zs5460", 8, 2, ""},
		{"zs5460", 2, 8, "5460"},
		{"zs5460", 0, 0, ""},
		{"", 2, 2, ""},
		{"中华人民共和国", 2, 2, "人民"},
		{"中华人民共和国", 8, 1, ""},
		{"公元2000年", 2, 5, "2000年"},
	}

	for _, test := range tests {
		if got := Mid(test.src, test.start, test.length); got != test.want {
			t.Errorf("Mid(%q,%d,%d) = %v, want %q",
				test.src,
				test.start,
				test.length,
				got,
				test.want)
		}
	}

}

func TestLen(t *testing.T) {
	var tests = []struct {
		src  string
		want int
	}{
		{"zs5460", 6},
		{"", 0},
		{"中华人民共和国", 7},
		{"公元2000年", 7},
	}

	for _, test := range tests {
		if got := Len(test.src); got != test.want {
			t.Errorf("Len(%q) = %v, want %v",
				test.src,
				got,
				test.want)
		}
	}
}
