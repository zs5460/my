package main

import (
	"github.com/zs5460/my"
)

func main() {
	println(my.Now())

	for i := 0; i < 10; i++ {
		println(my.RndStr(60))
	}

}
