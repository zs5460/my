package my

import (
	"strings"
	"testing"
)

func TestMD5(t *testing.T) {
	if MD5("zs5460") != "54c163ce691b59853ce7dc7df47028d7" {
		t.Error("MD5 value is wrong")
	}
}

func TestSHA1(t *testing.T) {
	if SHA1("zs5460") != "1fdab8d80dd4031533ddb77fbb62ffc56405f861" {
		t.Error("SHA1 value is wrong")
	}
}

func TestSHA256(t *testing.T) {
	if SHA256("zs5460") != "9f9430e4c60941321b24893d72086d771326a51bbd3bdd29cbcad229148631b4" {
		t.Error("SHA256 value is wrong")
	}
}

func TestHMACSHA1(t *testing.T) {
	if HMACSHA1("zs5460", "123456") != "9bd117811a0ff843655f0fbda7fc6fac404bffc4" {
		t.Error("HMAC_SHA1 value is wrong")
	}
}

func TestHMACSHA256(t *testing.T) {
	if HMACSHA256("zs5460", "123456") != "9b3900854cd8bb04e1c52de5ddfca7cde6b8e8179d16ef7d660b056d03e3609f" {
		t.Error("HMAC_SHA256 value is wrong")
	}
}

func BenchmarkHMACSHA1(b *testing.B) {
	s := strings.Repeat("hello", 1000)
	for i := 0; i < b.N; i++ {
		HMACSHA1(s, "123456")
	}
}
