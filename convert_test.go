package my

import (
	"strings"
	"testing"
)

func BenchmarkS2B(b *testing.B) {
	s := strings.Repeat("hello", 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b := str2bytes(s)
		_ = bytes2str(b)

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
