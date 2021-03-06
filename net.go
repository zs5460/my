package my

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	// RequestTimeout set the maximum time for a request.
	RequestTimeout time.Duration = 10
)

// GetURL request a url.
func GetURL(url string) (reply []byte, err error) {
	cli := &http.Client{
		Timeout: RequestTimeout * time.Second,
	}
	resp, err := cli.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	reply, err = io.ReadAll(resp.Body)
	return
}

// GetJSON get json from a url and unmarshal to a struct.
func GetJSON(url string, v interface{}) error {
	cli := &http.Client{
		Timeout: RequestTimeout * time.Second,
	}
	resp, err := cli.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	reply, _ := io.ReadAll(resp.Body)
	return json.Unmarshal(reply, v)
}

// PostURL request a url with POST method.
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
	reply, err = io.ReadAll(resp.Body)
	return
}

// PostJSON request a url with POST method
// params is a json string
func PostJSON(url string, params string) (reply []byte, err error) {
	cli := &http.Client{
		Timeout: RequestTimeout * time.Second,
	}
	resp, err := cli.Post(url,
		"application/json;charset=UTF-8",
		strings.NewReader(params))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	reply, err = io.ReadAll(resp.Body)
	return
}

// DownloadFile download a file from a url.
func DownloadFile(url, filepath string) (err error) {
	cli := &http.Client{}
	resp, err := cli.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	file, err := os.Create(filepath)
	if err != nil {
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return
}
