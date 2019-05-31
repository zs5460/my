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
