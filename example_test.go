package my_test

import (
	"fmt"

	"github.com/zs5460/my"
)

func ExampleGetJSON() {
	type Result struct {
		Code    int           `json:"code"`
		Message string        `json:"message"`
		Data    []interface{} `json:"data"`
		Total   int           `json:"total"`
	}
	var ret Result
	err := my.GetJSON("http://abc.com/api/.../", &ret)
	if err != nil {
		return
	}
	fmt.Println(ret.Code)
}
