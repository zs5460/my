package my_test

import (
	"fmt"
	"time"

	"github.com/zs5460/my"
)

func ExampleGetURL() {
	s, err := my.GetURL("http://54600.net/demo/ping")
	if err != nil {
		return
	}
	fmt.Println(string(s))
	// Output:
	// PONG
}

func ExampleGetJSON() {
	type Result struct {
		Code    int           `json:"code"`
		Message string        `json:"message"`
		Data    []interface{} `json:"data"`
		Total   int           `json:"total"`
	}
	var ret Result
	err := my.GetJSON("http://54600.net/demo/json/", &ret)
	if err != nil {
		return
	}
	fmt.Println(ret.Code)
	// Output:
	// 0
}

func ExamplePostURL() {
	s, err := my.PostURL("http://54600.net/demo/post/", "name=zs&age=20")
	if err != nil {
		return
	}
	fmt.Println(string(s))
	// Output:
	// name:zs
	// age:20
}

func ExampleFormatDateTime() {
	birthday := time.Date(1982, time.February, 11, 20, 13, 14, 0, time.Local)
	fmt.Println(my.FormatDateTime(birthday, 0))
	fmt.Println(my.FormatDateTime(birthday, 1))
	fmt.Println(my.FormatDateTime(birthday, 2))
	fmt.Println(my.FormatDateTime(birthday, 3))
	fmt.Println(my.FormatDateTime(birthday, 4))
	// Output:
	// 1982-02-11 20:13:14
	// 1982-02-11
	// 02-11
	// 20:13:14
	// 20:13
}

func ExampleLoadJSONConfig() {
	type config struct {
		AppName string
		Version string
	}
	var c config
	my.LoadJSONConfig("testdata/config.json", &c)
	fmt.Println(c.AppName)
	fmt.Println(c.Version)
	// Output:
	// DEMO
	// 1.0.0
}

func ExampleLoadXMLConfig() {
	type config struct {
		AppName string `xml:"appname"`
		Version string `xml:"version"`
	}
	var c config
	my.LoadXMLConfig("testdata/config.xml", &c)
	fmt.Println(c.AppName)
	fmt.Println(c.Version)
	// Output:
	// DEMO
	// 1.0.0
}
