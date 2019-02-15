package my

import (
	"strconv"
	"time"
)

// Now returns the current date and time.
func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// FriendlyTime ...
func FriendlyTime(t time.Time) string {
	dur := int(time.Now().Sub(t).Seconds())
	println(dur)
	switch {
	case dur < 60:
		return cstr(dur) + "刚刚"
	case dur < 3600:
		return cstr(dur/60) + "分钟前"
	case dur < 86400:
		return cstr(dur/3600) + "小时前"
	default:
		return t.Format("2006-01-02")
	}
}

func cstr(i int) string {
	return strconv.Itoa(i)
}
