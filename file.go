package my

import (
	"io/ioutil"
	"os"
)

// IsExist checks whether a file or directory exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// ReadText ...
func ReadText(filepath string) (text string, err error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}
	text = string(b)
	return
}

// WriteText ...
func WriteText(filepath string, text string) (err error) {
	return ioutil.WriteFile(filepath, []byte(text), 0666)

}
