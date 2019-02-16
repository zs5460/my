package main

import (
	"fmt"

	"github.com/zs5460/my"
)

func main() {
	s, _ := my.GetURL("http://localhost/ping")
	fmt.Printf("%s\n", s)
	s, _ = my.PostURL("http://localhost/post", "name=zs5460&age=20")
	fmt.Printf("%s\n", s)

}
