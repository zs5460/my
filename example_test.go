package my

import "fmt"

func ExampleGetJSON() {
	type Result struct {
		code    int
		message string
		data    []interface{}
	}
	var ret Result
	err := GetJSON("http://abc.com/api/.../", &ret)
	if err != nil {
		return
	}
	fmt.Println(ret.code)
}
