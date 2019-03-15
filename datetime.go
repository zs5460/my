package my

import (
	"time"
)

const (
	myLongDate  = 1
	myShortDate = 2
	myLongTime  = 3
	myShortTime = 4
)

// Now returns the current date and time.
func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// FriendlyTime returns easy to read format.
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

// FormatDateTime returns an expression formatted as a date or time.
func FormatDateTime(t time.Time, mode int) string {
	switch mode {
	case myLongDate:
		return t.Format("2006-01-02")
	case myShortDate:
		return t.Format("01-02")
	case myLongTime:
		return t.Format("15:04:05")
	case myShortTime:
		return t.Format("15:04")
	default:
		return t.Format("2006-01-02 15:04:05")
	}
}
