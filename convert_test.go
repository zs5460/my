package my

import (
	"strings"
	"testing"
	"time"
)

func BenchmarkS2B(b *testing.B) {
	s := strings.Repeat("hello", 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = StringToBytes(s)
	}
}
func BenchmarkB2S(b *testing.B) {
	bs := []byte(strings.Repeat("hello", 100))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = string(bs)
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

func TestCDate(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected time.Time
	}{
		{
			name:     "full datetime",
			input:    "2023-1-2 15:04:05",
			expected: time.Date(2023, 1, 2, 15, 4, 5, 0, time.Local),
		},
		{
			name:     "date only",
			input:    "2023-1-2",
			expected: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
		},
		{
			name:     "invalid format",
			input:    "invalid",
			expected: time.Date(1900, 1, 1, 0, 0, 0, 0, time.Local),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := CDate(tc.input)
			if !result.Equal(tc.expected) {
				t.Errorf("CDate(%q) = %v, want %v", tc.input, result, tc.expected)
			}
		})
	}
}
