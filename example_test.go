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

func ExampleCDate() {
	fmt.Println(my.CDate("1982-2-11 20:13:14"))
	fmt.Println(my.CDate("1982-02-11 20:13:14"))
	fmt.Println(my.CDate("1982-02-11"))
	fmt.Println(my.CDate("82-02-11")) //unsupport

	// Output:
	// 1982-02-11 20:13:14 +0800 CST
	// 1982-02-11 20:13:14 +0800 CST
	// 1982-02-11 00:00:00 +0800 CST
	// 1900-01-01 00:00:00 +0800 CST
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

func ExampleFriendlyTime() {
	fmt.Println(my.FriendlyTime(time.Now().Add(-2 * time.Second)))
	fmt.Println(my.FriendlyTime(time.Now().Add(-2 * time.Minute)))
	fmt.Println(my.FriendlyTime(time.Now().Add(-2 * time.Hour)))
	fmt.Println(my.FriendlyTime(time.Now().Add(-2 * 24 * time.Hour)))
	longlongago, _ := time.Parse("2006-01-02 15:04:05", "1982-02-11 20:13:14")
	fmt.Println(my.FriendlyTime(longlongago))

	// Output:
	// 刚刚
	// 2分钟前
	// 2小时前
	// 2天前
	// 1982-02-11
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

func ExampleFileExist() {
	fmt.Println(my.FileExist("testdata/config.json"))
	fmt.Println(my.FileExist("testdata/notexist.txt"))
	// Output:
	// true
	// false
}

func ExampleFolderExist() {
	fmt.Println(my.FolderExist("testdata"))
	fmt.Println(my.FolderExist("testdata2"))

	// Output:
	// true
	// false
}

func ExampleAppPath() {
	dir := my.AppPath()
	// dir is your app startup dir
	dir = "c:\\myapp"
	fmt.Println(dir)

	// Output:
	// c:\myapp
}
