package my

import (
	"os"
	"testing"
)

func TestGetURL(t *testing.T) {
	RequestTimeout = 2
	url := "https://www.nothissite.com/"
	c, err := GetURL(url)
	if err == nil {
		t.Fatalf("GetURL %s: error expected, none found", url)
	}
	url = "https://www.google.cn/"
	c, err = GetURL(url)
	if err != nil {
		t.Fatalf("GetURL %s: %v", url, err)
	}
	t.Log(c)
}

func TestGetJSON(t *testing.T) {

	type result struct {
		Code    int           `json:"code"`
		Message string        `json:"message"`
		Data    []interface{} `json:"data"`
		Total   int           `json:"total"`
	}

	RequestTimeout = 2
	var v = &result{}
	url := "https://www.nothissite.com/"
	err := GetJSON(url, &v)
	if err == nil {
		t.Fatalf("GetJSON %s: error expected, none found", url)
	}

	url = "https://www.google.cn/"
	err = GetJSON(url, &v)
	if err == nil {
		t.Fatalf("GetJSON %s: error expected, none found", url)
	}

	url = "http://54600.net/demo/json/"
	err = GetJSON(url, v)
	if err != nil {
		t.Fatalf("GetJSON %s: %v", url, err)
	}

	if v.Message != "ok" {
		t.Fatalf("message = %q\nexpected = %q", v.Message, "ok")
	}

}

func TestPostURL(t *testing.T) {
	url := "https://www.nothissite.com/"
	c, err := PostURL(url, "q=test")
	if err == nil {
		t.Fatalf("GetJSON %s: error expected, none found", url)
	}
	url = "http://54600.net/demo/post/"
	c, err = PostURL(url, "name=zs")
	if err != nil {
		t.Fatalf("PostURL %s: %v", url, err)
	}
	if bytes2str(c) != "name:zs\n" {
		t.Fatalf("body = %q\nexpected = %q", c, "name:zs\n")
	}
}

func TestDownloadFile(t *testing.T) {
	url := "https://www.nothissite.com/"
	localfile := "test.txt"
	err := DownloadFile(url, localfile)
	if err == nil {
		t.Fatalf("DownloadFile %s: error expected, none found", url)
	}

	url = "http://54600.net/demo/file/demo.txt"
	localfile = "thisdirnotexist/demo"
	err = DownloadFile(url, localfile) // create file error
	if err == nil {
		t.Fatalf("DownloadFile %s: error expected, none found", url)
	}

	url = "http://54600.net/demo/file/demo.txt"
	localfile = "testdata/demo.txt"
	err = DownloadFile(url, localfile)
	if err != nil {
		t.Fatalf("DownloadFile %s: %v", url, err)
	}

	if !FileExist(localfile) {
		t.Error("DownloadFile failed")
	}

	os.Remove(localfile) // clean up

}
