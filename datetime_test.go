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
