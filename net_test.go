package my_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	. "github.com/zs5460/my"
)

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.String()

		switch url {
		case "/ping":
			w.Header().Set("content-type", "text/plain")
			_, _ = fmt.Fprint(w, "PONG")

		case "/json":
			w.Header().Set("content-type", "application/json")
			_, _ = fmt.Fprint(w, `{"code":0,"message":"ok"}`)

		case "/post":
			w.Header().Set("content-type", "text/plain")
			_ = r.ParseForm()
			for k := range r.PostForm {
				_, _ = fmt.Fprint(w, k)
				_, _ = fmt.Fprint(w, ":")
				_, _ = fmt.Fprint(w, r.PostFormValue(k))
				_, _ = fmt.Fprint(w, "\n")
			}

		case "/postjson":
			w.Header().Set("content-type", "application/json")
			body,_ := ioutil.ReadAll(r.Body)
			_, _ = fmt.Fprint(w, `{"code":0,"data":`,string(body),`}`)

		case "/file/demo.txt":
			w.Header().Set("content-type", "text/plain")
			_, _ = fmt.Fprint(w, "this is a demo.")

		case "/404":
			w.WriteHeader(http.StatusNotFound)

		case "/500":
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.WriteHeader(500)
			_, _ = fmt.Fprintln(w, "internal server error")

		default:
			w.Header().Set("content-type", "text/plain")
			_, _ = fmt.Fprint(w, "hello")
		}

	}

	return httptest.NewServer(http.HandlerFunc(f))
}

func TestGetURL(t *testing.T) {
	ms := mockServer()
	defer ms.Close()

	url := "https://www.nothissite.com/"
	_, err := GetURL(url)
	if err == nil {
		t.Fatalf("GetURL %s: error expected, none found", url)
	}
	url = ms.URL + "/ping"
	_, err = GetURL(url)
	if err != nil {
		t.Fatalf("GetURL %s: %v", url, err)
	}
}

func TestGetJSON(t *testing.T) {
	ms := mockServer()
	defer ms.Close()

	type result struct {
		Code    int           `json:"code"`
		Message string        `json:"message"`
		Data    []interface{} `json:"data"`
		Total   int           `json:"total"`
	}

	var v = &result{}
	url := "https://www.nothissite.com/"
	err := GetJSON(url, &v)
	if err == nil {
		t.Fatalf("GetJSON %s: error expected, none found", url)
	}

	url = ms.URL + "/ping"
	err = GetJSON(url, &v)
	if err == nil {
		t.Fatalf("GetJSON %s: error expected, none found", url)
	}

	url = ms.URL + "/json"
	err = GetJSON(url, v)
	if err != nil {
		t.Fatalf("GetJSON %s: %v", url, err)
	}

	if v.Message != "ok" {
		t.Fatalf("message = %q\nexpected = %q", v.Message, "ok")
	}

}

func TestPostURL(t *testing.T) {
	ms := mockServer()
	defer ms.Close()

	url := "https://www.nothissite.com/"
	_, err := PostURL(url, "q=test")
	if err == nil {
		t.Fatalf("GetJSON %s: error expected, none found", url)
	}
	url = ms.URL + "/post"
	c, err := PostURL(url, "name=zs")
	if err != nil {
		t.Fatalf("PostURL %s: %v", url, err)
	}
	if BytesToString(c) != "name:zs\n" {
		t.Fatalf("body = %q\nexpected = %q", c, "name:zs\n")
	}
}

func TestPostJSON(t *testing.T) {
	ms := mockServer()
	defer ms.Close()

	url := ms.URL + "/postjson"
	c, err := PostJSON(url, `{"name":"zs"}`)
	if err != nil {
		t.Fatalf("PostJSON %s: %v", url, err)
	}
	if BytesToString(c) != `{"code":0,"data":{"name":"zs"}}` {
		t.Fatalf("body = %q\nexpected = %q", c, "name:zs\n")
	}
}


func TestDownloadFile(t *testing.T) {
	ms := mockServer()
	defer ms.Close()

	url := "https://www.nothissite.com/"
	localfile := "test.txt"
	err := DownloadFile(url, localfile)
	if err == nil {
		t.Fatalf("DownloadFile %s: error expected, none found", url)
	}

	url = ms.URL + "/file/demo.txt"
	localfile = "thisdirnotexist/demo"
	err = DownloadFile(url, localfile) // create file error
	if err == nil {
		t.Fatalf("DownloadFile %s: error expected, none found", url)
	}

	localfile = "testdata/demo.txt"
	err = DownloadFile(url, localfile)
	if err != nil {
		t.Fatalf("DownloadFile %s: %v", url, err)
	}

	if !FileExist(localfile) {
		t.Error("DownloadFile failed")
	}

	_ = os.Remove(localfile) // clean up
}
