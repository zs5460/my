package my

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
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
	resp, err := requestClient().Get(url)
	if err != nil {
		return
	}
	defer func() { _ = resp.Body.Close() }()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	reply, err = io.ReadAll(resp.Body)
	return
}

// GetJSON get json from a url and unmarshal to a struct.
func GetJSON(url string, v any) error {
	resp, err := requestClient().Get(url)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if err = checkResponse(resp); err != nil {
		return err
	}
	reply, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(reply, v)
}

// PostURL request a url with POST method.
// params is a string like k1=v1&k2=v2
func PostURL(url string, params string) (reply []byte, err error) {
	resp, err := requestClient().Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(params))
	if err != nil {
		return
	}
	defer func() { _ = resp.Body.Close() }()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	reply, err = io.ReadAll(resp.Body)
	return
}

// PostJSON request a url with POST method
// params is a json string
func PostJSON(url string, params string) (reply []byte, err error) {
	resp, err := requestClient().Post(url,
		"application/json;charset=UTF-8",
		strings.NewReader(params))
	if err != nil {
		return
	}
	defer func() { _ = resp.Body.Close() }()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	reply, err = io.ReadAll(resp.Body)
	return
}

// DownloadFile download a file from a url.
func DownloadFile(url, filepath string) (err error) {
	resp, err := (&http.Client{}).Get(url)
	if err != nil {
		return
	}
	defer func() { _ = resp.Body.Close() }()
	if err = checkResponse(resp); err != nil {
		return err
	}

	file, err := os.Create(filepath)
	if err != nil {
		return
	}
	defer func() { _ = file.Close() }()

	_, err = io.Copy(file, resp.Body)
	return
}

func requestClient() *http.Client {
	return &http.Client{Timeout: RequestTimeout * time.Second}
}

func checkResponse(resp *http.Response) error {
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("http request failed: %s", resp.Status)
	}
	return nil
}

// Handler wraps the http.Handler with panic recovery support.
func RecoverWrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			}
		}()

		h.ServeHTTP(w, r)
	})
}

// RemoteIP
func RemoteIP(r *http.Request) string {
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

// ClientIP
func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
