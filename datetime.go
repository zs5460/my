package my

import "time"

// Now ...
func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
