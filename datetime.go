package my

import (
	"time"
)

// Now returns the current date and time.
func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// FriendlyTime ...
func FriendlyTime(t time.Time) string {
	seconds := int(time.Now().Sub(t).Seconds())
	switch {
	case seconds > 0 && seconds < 60:
		return "刚刚"
	case seconds >= 60 && seconds < 3600:
		return CStr(seconds/60) + "分钟前"
	case seconds >= 3600 && seconds < 86400:
		return CStr(seconds/3600) + "小时前"
	case seconds >= 86400 && seconds < 604800:
		return CStr(seconds/86400) + "天前"
	default:
		return t.Format("2006-01-02")
	}
}
