package my

import (
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
