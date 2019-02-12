package my

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

// MD5 ...
func MD5(s string) (ret string) {
	h := md5.New()
	io.WriteString(h, s)
	ret = hex.EncodeToString(h.Sum(nil))
	return
}

// SHA1 ...
func SHA1(s string) (ret string) {
	h := sha1.New()
	io.WriteString(h, s)
	ret = hex.EncodeToString(h.Sum(nil))
	return
}

// SHA256 ...
func SHA256(s string) (ret string) {
	h := sha256.New()
	io.WriteString(h, s)
	ret = hex.EncodeToString(h.Sum(nil))
	return
}
