package my

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
)

// LoadJSONConfig load config from json file.
func LoadJSONConfig(fn string, v interface{}) {
	content, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, v)
	if err != nil {
		log.Fatal(err)
	}

}

// LoadXMLConfig load config from xml file.
func LoadXMLConfig(fn string, v interface{}) {
	content, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}
	err = xml.Unmarshal(content, v)
	if err != nil {
		log.Fatal(err)
	}
}
