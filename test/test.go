package main

import (
	"fmt"
	"github.com/zs5460/my"
)

func main() {
	/*
		s, _ := my.GetURL("http://localhost/ping")
		fmt.Printf("%s\n", s)
		s, _ = my.PostURL("http://localhost/post", "name=zs5460&age=20")
		fmt.Printf("%s\n", s)


			my.RequestTimeout = 5
			s, err := my.GetURL("http://localhost/longtime")
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			fmt.Printf("%s\n", s)
	*/
	d, _ := my.CDate("2019-02-18 14:40:00")

	//fmt.Println(d)
	fmt.Println(my.FriendlyTime(d))
	fmt.Println(my.Now())

	fmt.Println("done!")
}
