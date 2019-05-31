package my_test

import (
	"strings"
	"testing"

	. "github.com/zs5460/my"
)

func BenchmarkS2B(b *testing.B) {
	s := strings.Repeat("hello", 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b := StringToBytes(s)
		_ = BytesToString(b)

	}
}
func BenchmarkS2B0(b *testing.B) {
	s := strings.Repeat("hello", 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b := []byte(s)
		_ = string(b)

	}
}

func TestCStr(t *testing.T) {
	if CStr(0) != "0" {
		t.Errorf("CStr did not work properly")
	}
}

func TestCInt(t *testing.T) {
	testCases := []struct {
		s string
		i int
	}{
		{"100", 100},
		{"1", 1},
		{"0", 0},
		{"abcd", 0},
	}
	for _, tC := range testCases {
		t.Run(tC.s, func(t *testing.T) {
			if got := CInt(tC.s); got != tC.i {
				t.Errorf("CInt(%q) want %d got %d", tC.s, tC.i, got)
			}
		})
	}
}

func TestStringToBytes(t *testing.T) {
	s := "hello"
	ss := []byte{'h', 'e', 'l', 'l', 'o'}
	if BytesToString(ss) != s {
		t.Errorf("BytesToString did not work properly")
	}

	if StringToBytes(s)[0] != 'h' {
		t.Errorf("StringToBytes did not work properly")
	}

	if BytesToString(StringToBytes(s)) != s {
		t.Errorf("StringToBytes and BytesToString did not work properly")
	}

}
