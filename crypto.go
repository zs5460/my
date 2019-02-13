package my

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

// MD5 ...
func MD5(s string) (ret string) {
	var h hash.Hash
	h = md5.New()
	h.Write([]byte(s))
	ret = hex.EncodeToString(h.Sum(nil))
	return
}

// SHA1 ...
func SHA1(s string) (ret string) {
	var h hash.Hash
	h = sha1.New()
	h.Write([]byte(s))
	ret = hex.EncodeToString(h.Sum(nil))
	return
}

// SHA256 ...
func SHA256(s string) (ret string) {
	var h hash.Hash
	h = sha256.New()
	h.Write([]byte(s))
	ret = hex.EncodeToString(h.Sum(nil))
	return
}

// HMACSHA1 ...
func HMACSHA1(s string, key string) (ret string) {
	var hashFunc = sha1.New
	h := hmac.New(hashFunc, []byte(key))
	h.Write([]byte(s))
	ret = hex.EncodeToString(h.Sum(nil))
	return
}
