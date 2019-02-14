package my

import "testing"

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
