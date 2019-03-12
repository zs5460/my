package my_test

import (
	"fmt"

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
