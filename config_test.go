package my_test

import (
	"testing"

	. "github.com/zs5460/my"
)

type config struct {
	AppName string `xml:"appname"`
	Version string `xml:"version"`
}

func TestLoadJSONConfig(t *testing.T) {
	var c config
	cfg := "testdata/noexist.json"
	err := LoadJSONConfig(cfg, &c)
	if err == nil {
		t.Fatalf("LoadJSONConfig %s: error expected, none found", cfg)
	}

	cfg = "testdata/config.json"
	err = LoadJSONConfig(cfg, &c)
	if err != nil {
		t.Fatalf("LoadJSONConfig %s: %v", cfg, err)
	}

	if c.AppName != "DEMO" || c.Version != "1.0.0" {
		t.Error("load json config failed!")
	}

}

func TestLoadXMLConfig(t *testing.T) {
	var c config
	cfg := "testdata/noexist.xml"
	err := LoadXMLConfig(cfg, &c)
	if err == nil {
		t.Fatalf("LoadXMLConfig %s: error expected, none found", cfg)
	}

	cfg = "testdata/config.xml"
	err = LoadXMLConfig(cfg, &c)
	if err != nil {
		t.Fatalf("LoadXMLConfig %s: %v", cfg, err)
	}

	if c.AppName != "DEMO" || c.Version != "1.0.0" {
		t.Error("load xml config failed!")
	}

}
