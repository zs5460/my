package my_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/zs5460/my"
)

func ExampleLoadJSONConfig() {
	type config struct {
		AppName string
		Version string
	}
	var c config
	err := my.LoadJSONConfig("testdata/config.json", &c)
	if err != nil {
		log.Fatal(err)
	}
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
	err := my.LoadXMLConfig("testdata/config.xml", &c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c.AppName)
	fmt.Println(c.Version)
	// Output:
	// DEMO
	// 1.0.0
}

func ExampleMustLoadConfig() {
	type config struct {
		AppName string `xml:"appname"`
		Version string `xml:"version"`
	}
	var c config
	my.MustLoadConfig("testdata/config.xml", &c)
	fmt.Println(c.AppName)
	fmt.Println(c.Version)
	// Output:
	// DEMO
	// 1.0.0
}

func ExampleIsEmpty() {
	var s string
	fmt.Println(my.IsEmpty(s))
	s = "zs5460"
	fmt.Println(my.IsEmpty(s))
	// Output:
	// true
	// false
}

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

func ExampleDownloadFile() {
	source := "https://avatars3.githubusercontent.com/u/1711957"
	localfile := "testdata/avatar.jpg"
	err := my.DownloadFile(source, localfile)
	if err != nil {
		log.Println(err)
	}
	os.Remove(localfile) //clean up
	fmt.Println("OK")
	// Output:
	// OK
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

func ExampleLeft() {
	s := "zs5460"
	fmt.Println(my.Left(s, 2))
	s = "公元2000年"
	fmt.Println(my.Left(s, 5))
	// Output:
	// zs
	// 公元200
}

func ExampleRight() {
	s := "zs5460"
	fmt.Println(my.Right(s, 2))
	s = "公元2000年"
	fmt.Println(my.Right(s, 5))
	// Output:
	// 60
	// 2000年
}

func ExampleMid() {
	s := "zs5460"
	fmt.Println(my.Mid(s, 2, 2))
	s = "公元2000年"
	fmt.Println(my.Mid(s, 2, 4))
	// Output:
	// 54
	// 2000
}

func ExampleLen() {
	s := "zs5460"
	fmt.Println(my.Len(s))
	s = "公元2000年"
	fmt.Println(my.Len(s))
	// Output:
	// 6
	// 7
}

func ExampleSpace() {
	s := "a" + my.Space(5) + "b"
	fmt.Println(s)
	fmt.Println("1234567")
	// Output:
	// a     b
	// 1234567
}

func ExampleTest() {
	// pattern can be english,chinese,idcard,email,qq,zip,phone,mobile,url,username,password.
	s := "zs5460"
	fmt.Println(my.Test(s, "english"))
	s = "zhousong"
	fmt.Println(my.Test(s, "english"))
	s = "zs5460.gmail.com"
	fmt.Println(my.Test(s, "email"))
	s = "zs5460@gmail.com"
	fmt.Println(my.Test(s, "email"))

	// Output:
	// false
	// true
	// false
	// true
}
