package my

import (
	"os"
)

// IsExist ...
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
