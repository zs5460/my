package my

import (
	"time"
)

type dateTimeFormat int

const (
	myDateTime dateTimeFormat = iota
	myLongDate
	myShortDate
	myLongTime
	myShortTime
)

// Now returns the current date and time formatted as "2006-01-02 15:04:05".
func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// FriendlyTime returns a human-friendly representation of the time elapsed since the given time.
// It returns different formats based on the elapsed time:
// - "刚刚" for less than a minute
// - "X分钟前" for less than an hour
// - "X小时前" for less than a day
// - "X天前" for less than a week
// - The date in "2006-01-02" format for longer periods
func FriendlyTime(t time.Time) string {
	seconds := int(time.Since(t).Seconds())
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

// FormatDateTime formats a time.Time according to the specified mode.
// Available modes are:
// - myDateTime: "2006-01-02 15:04:05"
// - myLongDate: "2006-01-02"
// - myShortDate: "01-02"
// - myLongTime: "15:04:05"
// - myShortTime: "15:04"
func FormatDateTime(t time.Time, mode dateTimeFormat) string {
	switch mode {
	case myDateTime:
		return t.Format("2006-01-02 15:04:05")
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
