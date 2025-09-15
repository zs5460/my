package my

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	now := time.Now().Format("2006-01-02 15:04:05")
	if Now() != now {
		t.Errorf("Now() did not work properly")
	}
}

func TestFormatDateTime(t *testing.T) {
	testTime := time.Date(2023, 5, 15, 14, 30, 45, 0, time.Local)
	testCases := []struct {
		name     string
		mode     dateTimeFormat
		expected string
	}{
		{"DateTime", myDateTime, "2023-05-15 14:30:45"},
		{"LongDate", myLongDate, "2023-05-15"},
		{"ShortDate", myShortDate, "05-15"},
		{"LongTime", myLongTime, "14:30:45"},
		{"ShortTime", myShortTime, "14:30"},
		{"Default", dateTimeFormat(99), "2023-05-15 14:30:45"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := FormatDateTime(testTime, tc.mode)
			if result != tc.expected {
				t.Errorf("FormatDateTime with mode %v = %q, want %q", tc.mode, result, tc.expected)
			}
		})
	}
}

func TestFriendlyTime(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{"Just now", now.Add(-30 * time.Second), "刚刚"},
		{"Minutes ago", now.Add(-10 * time.Minute), "10分钟前"},
		{"Hours ago", now.Add(-3 * time.Hour), "3小时前"},
		{"Days ago", now.Add(-2 * 24 * time.Hour), "2天前"},
		{"Long time ago", now.Add(-10 * 24 * time.Hour), now.Add(-10 * 24 * time.Hour).Format("2006-01-02")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := FriendlyTime(tc.input)
			if tc.name != "Long time ago" && result != tc.expected {
				t.Errorf("FriendlyTime(%v) = %q, want %q", tc.input, result, tc.expected)
			}
		})
	}
}
