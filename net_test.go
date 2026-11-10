package my

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
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
				_, _ = fmt.Fprint(w, k, ":", r.PostFormValue(k), "\n")
			}

		case "/postjson":
			w.Header().Set("content-type", "application/json")
			body, _ := io.ReadAll(r.Body)
			_, _ = fmt.Fprint(w, `{"code":0,"data":`, string(body), `}`)

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
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    []any  `json:"data"`
		Total   int    `json:"total"`
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

	_, err = GetURL(ms.URL + "/404")
	if err == nil {
		t.Fatal("GetURL should reject non-2xx responses")
	}
}

func TestPostURL(t *testing.T) {
	ms := mockServer()
	defer ms.Close()

	url := "https://www.nothissite.com/"
	_, err := PostURL(url, "q=test")
	if err == nil {
		t.Fatalf("PostURL %s: error expected, none found", url)
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
	url := "https://www.nothissite.com/"
	_, err := PostJSON(url, `{"name":"zs"}`)
	if err == nil {
		t.Fatalf("PostJSON %s: error expected, none found", url)
	}

	url = ms.URL + "/postjson"
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
	localFile := "testdata/demo.txt"
	err := DownloadFile(url, localFile)
	if err == nil {
		t.Fatalf("DownloadFile %s: error expected, none found", url)
	}

	url = ms.URL + "/file/demo.txt"
	localFile = "this_dir_not_exist/demo"
	err = DownloadFile(url, localFile) // create file error
	if err == nil {
		t.Fatalf("DownloadFile %s: error expected, none found", url)
	}

	localFile = "testdata/demo.txt"
	err = DownloadFile(url, localFile)
	if err != nil {
		t.Fatalf("DownloadFile %s: %v", url, err)
	}

	if !FileExist(localFile) {
		t.Error("DownloadFile failed")
	}

	_ = os.Remove(localFile) // clean up
}

func captureLog(t *testing.T) *bytes.Buffer {
	t.Helper()
	var buf bytes.Buffer
	old := log.Writer()
	log.SetOutput(&buf)
	t.Cleanup(func() { log.SetOutput(old) })
	return &buf
}

func TestRecoverWrap(t *testing.T) {
	logs := captureLog(t)
	h := RecoverWrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("boom")
	}))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusInternalServerError)
	}
	if body, want := rec.Body.String(), http.StatusText(http.StatusInternalServerError); body != want {
		t.Fatalf("body = %q, want %q", body, want)
	}
	if !strings.Contains(logs.String(), "panic: boom") {
		t.Fatalf("panic was not logged: %q", logs.String())
	}
}

func TestRecoverWrapPassThrough(t *testing.T) {
	h := RecoverWrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
		_, _ = w.Write([]byte("ok"))
	}))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))
	if rec.Code != http.StatusTeapot || rec.Body.String() != "ok" {
		t.Fatalf("status = %d, body = %q", rec.Code, rec.Body.String())
	}
}

func TestRemoteIP(t *testing.T) {
	tests := []struct {
		addr string
		want string
	}{
		{"1.2.3.4:5678", "1.2.3.4"},
		{"[2001:db8::1]:443", "2001:db8::1"},
		{"1.2.3.4", ""},
		{"", ""},
	}
	for _, tt := range tests {
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		r.RemoteAddr = tt.addr
		if got := RemoteIP(r); got != tt.want {
			t.Errorf("RemoteIP(%q) = %q, want %q", tt.addr, got, tt.want)
		}
	}
}

func TestClientIP(t *testing.T) {
	tests := []struct {
		name   string
		header string
		value  string
		remote string
		want   string
	}{
		{"xff", "X-Forwarded-For", "9.9.9.9", "1.2.3.4:5678", "9.9.9.9"},
		{"xff list", "X-Forwarded-For", "9.9.9.9, 8.8.8.8", "1.2.3.4:5678", "9.9.9.9"},
		{"xff spaces", "X-Forwarded-For", "  9.9.9.9  ", "1.2.3.4:5678", "9.9.9.9"},
		{"xff empty", "X-Forwarded-For", "", "1.2.3.4:5678", "1.2.3.4"},
		{"x-real-ip", "X-Real-Ip", "7.7.7.7", "1.2.3.4:5678", "7.7.7.7"},
		{"remote", "", "", "1.2.3.4:5678", "1.2.3.4"},
		{"empty", "", "", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			if tt.header != "" {
				r.Header.Set(tt.header, tt.value)
			}
			r.RemoteAddr = tt.remote
			if got := ClientIP(r); got != tt.want {
				t.Errorf("ClientIP() = %q, want %q", got, tt.want)
			}
		})
	}
}
