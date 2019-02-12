package my

import "testing"

func TestLeft(t *testing.T) {
	if Left("zs5460", 2) != "zs" {
		t.Error("Left value is wrong")
	}
}

func TestRight(t *testing.T) {

	var tests = []struct {
		src    string
		lenght int
		want   string
	}{
		{"zs5460", 2, "60"},
		{"zs5460", 8, "zs5460"},
		{"zs5460", 1, "0"},
		{"zs5460", 0, "zs5460"},
		{"", 2, ""},
	}

	for _, test := range tests {
		if got := Right(test.src, test.lenght); got != test.want {
			t.Errorf("Right(%q,%d) = %v, want %s",
				test.src,
				test.lenght,
				got,
				test.want)
		}
	}

}
