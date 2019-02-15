package main

import (
	"fmt"
	"time"

	"github.com/zs5460/my"
)

func main() {
	t := time.Date(2019, time.February, 16, 15, 1, 1, 0, time.Local)
	fmt.Printf("%v\n", time.Now())
	fmt.Printf("%v\n", t)
	fmt.Printf("%v\n", my.FriendlyTime(t))

}
