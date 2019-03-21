package my

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// AppPath returns an absolute path of app.
func AppPath() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}

// FolderExist returns true if a specified folder exists; false if it does not.
func FolderExist(path string) bool {
	info, err := os.Stat(path)
	if err != nil || !info.IsDir() {
		return false
	}
	return true
}

// FileExist returns true if a specified file exists; false if it does not.
func FileExist(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// MakeDir creates a folder.
func MakeDir(dir string) error {
	if !FolderExist(dir) {
		err := os.MkdirAll(dir, os.ModePerm)
		return err
	}
	return nil
}

// ReadText reads an entire text file and returns the resulting string.
func ReadText(filepath string) (text string, err error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}
	text = string(b)
	return
}

// WriteText writes a specified string to a text file.
func WriteText(filepath string, text string) (err error) {
	//return ioutil.WriteFile(filepath, []byte(text), 0644)
	return writeText(filepath, text, false)
}

// AppendText append a string to a text file.
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
