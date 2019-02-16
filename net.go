package my

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// GetURL ...
func GetURL(url string) (reply []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	reply, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

// PostURL ...
// params is a string like k1=v1&k2=v2
func PostURL(url string, params string) (reply []byte, err error) {
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(params))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	reply, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
