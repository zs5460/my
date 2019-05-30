package my

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"path"
)

var configParser = make(map[string]func(string, interface{}) error, 0)

func init() {
	configParser["default"] = LoadJSONConfig
	configParser[".json"] = LoadJSONConfig
	configParser[".xml"] = LoadXMLConfig
}

// LoadJSONConfig load config from json file.
func LoadJSONConfig(fn string, v interface{}) error {
	content, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, v)
	return err
}

// LoadXMLConfig load config from xml file.
func LoadXMLConfig(fn string, v interface{}) error {
	content, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(content, v)
	return err
}

// MustLoadConfig load config or panic.
func MustLoadConfig(fn string, v interface{}) {
	configtype := path.Ext(fn)
	parser, exist := configParser[configtype]
	if !exist {
		panic("unsupport config file.")
	}
	err := parser(fn, v)
	if err != nil {
		panic("config file parse error.")
	}
}
