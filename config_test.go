package my_test

import (
	"testing"

	"github.com/zs5460/my"
)

type config struct {
	AppName string `xml:"appname"`
	Version string `xml:"version"`
}

func TestLoadJSONConfig(t *testing.T) {
	var c config
	my.LoadJSONConfig("testdata/config.json", &c)
	if c.AppName != "DEMO" || c.Version != "1.0.0" {
		t.Error("load json config failed!")
	}

}

func TestLoadXMLConfig(t *testing.T) {
	var c config
	my.LoadXMLConfig("testdata/config.xml", &c)
	if c.AppName != "DEMO" || c.Version != "1.0.0" {
		t.Error("load xml config failed!")
	}

}
