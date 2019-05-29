package my

import "testing"

func TestGetURL(t *testing.T) {
	RequestTimeout = 2
	c, err := GetURL("https://www.nothissite.com/")
	if err != nil {
		t.Log(err)
	}
	t.Log(c)
	c, err = GetURL("https://www.google.cn/")
	if err != nil {
		t.Log(err)
	}
	t.Log(c)
}
