package my

import (
	"fmt"
	"log"
	"os"
	"time"
)

func init() {
	RequestTimeout = 1
}

func ExampleLoadJSONConfig() {
	type config struct {
		AppName string
		Version string
	}
	var c config
	err := LoadJSONConfig("testdata/config.json", &c)
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
	err := LoadXMLConfig("testdata/config.xml", &c)
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
	MustLoadConfig("testdata/config.xml", &c)
	fmt.Println(c.AppName)
	fmt.Println(c.Version)
	// Output:
	// DEMO
	// 1.0.0
}

func ExampleCDate() {
	fmt.Println(CDate("1982-2-11 20:13:14").Format("2006-01-02 15:04:05"))
	fmt.Println(CDate("1982-02-11 20:13:14").Format("2006-01-02 15:04:05"))
	fmt.Println(CDate("1982-02-11").Format("2006-01-02 15:04:05"))
	fmt.Println(CDate("82-02-11").Format("2006-01-02 15:04:05")) //unsupport

	// Output:
	// 1982-02-11 20:13:14
	// 1982-02-11 20:13:14
	// 1982-02-11 00:00:00
	// 1900-01-01 00:00:00
}

func ExampleIsEmpty() {
	var s string
	fmt.Println(IsEmpty(s))
	s = "zs5460"
	fmt.Println(IsEmpty(s))
	// Output:
	// true
	// false
}

func ExampleFormatDateTime() {
	birthday := time.Date(1982, time.February, 11, 20, 13, 14, 0, time.Local)
	fmt.Println(FormatDateTime(birthday, 0))
	fmt.Println(FormatDateTime(birthday, 1))
	fmt.Println(FormatDateTime(birthday, 2))
	fmt.Println(FormatDateTime(birthday, 3))
	fmt.Println(FormatDateTime(birthday, 4))
	fmt.Println(FormatDateTime(birthday, 5))

	// Output:
	// 1982-02-11 20:13:14
	// 1982-02-11
	// 02-11
	// 20:13:14
	// 20:13
	// 1982-02-11 20:13:14
}

func ExampleFriendlyTime() {
	fmt.Println(FriendlyTime(time.Now().Add(-2 * time.Second)))
	fmt.Println(FriendlyTime(time.Now().Add(-2 * time.Minute)))
	fmt.Println(FriendlyTime(time.Now().Add(-2 * time.Hour)))
	fmt.Println(FriendlyTime(time.Now().Add(-2 * 24 * time.Hour)))
	longLongAgo, _ := time.Parse("2006-01-02 15:04:05", "1982-02-11 20:13:14")
	fmt.Println(FriendlyTime(longLongAgo))

	// Output:
	// 刚刚
	// 2分钟前
	// 2小时前
	// 2天前
	// 1982-02-11
}

func ExampleLeft() {
	s := "zs5460"
	fmt.Println(Left(s, 2))
	s = "公元2000年"
	fmt.Println(Left(s, 5))
	// Output:
	// zs
	// 公元200
}

func ExampleRight() {
	s := "zs5460"
	fmt.Println(Right(s, 2))
	s = "公元2000年"
	fmt.Println(Right(s, 5))
	// Output:
	// 60
	// 2000年
}

func ExampleMid() {
	s := "zs5460"
	fmt.Println(Mid(s, 2, 2))
	s = "公元2000年"
	fmt.Println(Mid(s, 2, 4))
	// Output:
	// 54
	// 2000
}

func ExampleLen() {
	s := "zs5460"
	fmt.Println(Len(s))
	s = "公元2000年"
	fmt.Println(Len(s))
	// Output:
	// 6
	// 7
}

func ExampleSpace() {
	s := "a" + Space(5) + "b"
	fmt.Println(s)
	fmt.Println("1234567")
	// Output:
	// a     b
	// 1234567
}

func ExampleTest() {
	// pattern can be english,chinese,idcard,email,qq,zip,phone,mobile,url,username,password.
	s := "zs5460"
	fmt.Println(Test(s, "english"))
	s = "zhousong"
	fmt.Println(Test(s, "english"))
	s = "zs5460.gmail.com"
	fmt.Println(Test(s, "email"))
	s = "zs5460@gmail.com"
	fmt.Println(Test(s, "email"))

	// Output:
	// false
	// true
	// false
	// true
}

func ExampleFileExist() {
	fmt.Println(FileExist("testdata/config.json"))
	fmt.Println(FileExist("testdata/notexist.txt"))
	// Output:
	// true
	// false
}

func ExampleFolderExist() {
	fmt.Println(FolderExist("testdata"))
	fmt.Println(FolderExist("testdata2"))

	// Output:
	// true
	// false
}

func ExampleAppPath() {
	dir := AppPath()
	// dir is your app startup dir
	dir = "c:\\fakeapp"
	fmt.Println(dir)

	// Output:
	// c:\fakeapp
}

func ExampleGetURL() {
	ms := mockServer()
	defer ms.Close()
	s, err := GetURL(ms.URL + "/ping")
	if err != nil {
		return
	}
	fmt.Println(string(s))
	// Output:
	// PONG
}

func ExampleGetJSON() {
	ms := mockServer()
	defer ms.Close()
	type Result struct {
		Code    int           `json:"code"`
		Message string        `json:"message"`
		Data    []interface{} `json:"data"`
		Total   int           `json:"total"`
	}
	var ret Result
	err := GetJSON(ms.URL+"/json", &ret)
	if err != nil {
		return
	}
	fmt.Println(ret.Code)
	// Output:
	// 0
}

func ExamplePostURL() {
	ms := mockServer()
	defer ms.Close()
	s, err := PostURL(ms.URL+"/post", "name=zs")
	if err != nil {
		return
	}
	fmt.Println(string(s))
	// Output:
	// name:zs
}

func ExampleDownloadFile() {
	ms := mockServer()
	defer ms.Close()
	localFile := "testdata/demo.txt"
	err := DownloadFile(ms.URL+"/file/demo.txt", localFile)
	if err != nil {
		log.Println(err)
	}
	os.Remove(localFile) //clean up
	fmt.Println("OK")
	// Output:
	// OK
}
