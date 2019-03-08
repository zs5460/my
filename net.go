package my

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var (
	// RequestTimeout ...
	RequestTimeout time.Duration = 10
)

// GetURL ...
func GetURL(url string) (reply []byte, err error) {
	cli := &http.Client{
		Timeout: RequestTimeout * time.Second,
	}
	resp, err := cli.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	reply, err = ioutil.ReadAll(resp.Body)
	return
}

// PostURL ...
// params is a string like k1=v1&k2=v2
func PostURL(url string, params string) (reply []byte, err error) {
	cli := &http.Client{
		Timeout: RequestTimeout * time.Second,
	}
	resp, err := cli.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(params))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	reply, err = ioutil.ReadAll(resp.Body)
	return
}

// DownloadFile .....
func DownloadFile(url, filepath string) error {
	cli := &http.Client{}
	resp, err := cli.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	return ioutil.WriteFile(filepath, data, 0666)

}
