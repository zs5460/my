package my

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

// MD5 calculate the md5 hash of a string.
func MD5(s string) string {
	var h hash.Hash
	h = md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))

}

// SHA1 calculate the sha1 hash of a string.
func SHA1(s string) string {
	var h hash.Hash
	h = sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// SHA256 calculate the sha256 hash of a string.
func SHA256(s string) string {
	var h hash.Hash
	h = sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// HMACSHA1 generate a keyed hash value using the HMAC method.
func HMACSHA1(s, key string) string {
	var hashFunc = sha1.New
	h := hmac.New(hashFunc, []byte(key))
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// HMACSHA256 generate a keyed hash value using the HMAC method.
func HMACSHA256(s, key string) string {
	var hashFunc = sha256.New
	h := hmac.New(hashFunc, []byte(key))
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
