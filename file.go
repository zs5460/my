package my

import (
	"io"
	"io/ioutil"
	"os"
)

// IsExist checks whether a file or directory exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// MakeDir ...
func MakeDir(dir string) error {
	if !IsExist(dir) {
		err := os.MkdirAll(dir, os.ModePerm)
		return err
	}
	return nil
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
	//return ioutil.WriteFile(filepath, []byte(text), 0644)
	return writeText(filepath, text, false)
}

// AppendText ...
func AppendText(filepath string, text string) (err error) {
	return writeText(filepath, text, true)
}

func writeText(filepath string, text string, appendMode bool) (err error) {
	var mode int
	if appendMode {
		mode = os.O_RDWR | os.O_CREATE | os.O_APPEND
	} else {
		mode = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	}
	f, err := os.OpenFile(filepath, mode, 0644)
	defer f.Close()
	if err != nil {
		return
	}
	io.WriteString(f, text)
	return
}
