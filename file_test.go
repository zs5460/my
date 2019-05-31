package my_test

import (
	"os"
	"testing"

	. "github.com/zs5460/my"
)

func TestFolderExist(t *testing.T) {
	if isexist := FolderExist("testdata"); !isexist {
		t.Errorf("expected %v, got %v", true, isexist)
	}
	if isexist := FolderExist("testdata2"); isexist {
		t.Errorf("expected %v, got %v", false, isexist)
	}
}

func TestFileExist(t *testing.T) {
	if isexist := FileExist("testdata/config.json"); !isexist {
		t.Errorf("expected %v, got %v", true, isexist)
	}
	if isexist := FileExist("testdata/jpm.txt"); isexist {
		t.Errorf("expected %v, got %v", false, isexist)
	}
}

func TestMakeDir(t *testing.T) {
	testdir := "testdata/subdir"
	err := MakeDir(testdir)
	if err != nil {
		t.Errorf("MakeDir did not work properly, %v", err)
	}
	err = MakeDir(testdir)
	if err != nil {
		t.Errorf("MakeDir did not work properly, %v", err)
	}
	os.Remove(testdir) //clean up
}

func TestReadText(t *testing.T) {
	_, err := ReadText("testdata/hello2")
	if err == nil {
		t.Errorf("ReadText did not work properly")
	}
	c, err := ReadText("testdata/hello")
	if err != nil {
		t.Fatal(err)
	}
	if c != "world" {
		t.Errorf("ReadText did not work properly")
	}
}

func TestWriteText(t *testing.T) {
	err := WriteText("testdata2/w.txt", "OK")
	if err == nil {
		t.Errorf("WriteText did not work properly")
	}
	fn := "testdata/w.txt"
	text := "this is a test"
	err = WriteText(fn, text)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(fn) // clean up
	c, err := ReadText(fn)
	if err != nil {
		t.Errorf("WriteText did not work properly")
	}
	if c != text {
		t.Errorf("WriteText did not work properly")
	}
}

func TestAppendText(t *testing.T) {
	fn := "testdata/w.txt"
	text := "test"
	err := AppendText(fn, text)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(fn) // clean up
	err = AppendText(fn, text)
	if err != nil {
		t.Fatal(err)
	}
	c, err := ReadText(fn)
	if err != nil {
		t.Errorf("AppendText did not work properly")
	}
	if c != "testtest" {
		t.Errorf("AppendText did not work properly")
	}
}
