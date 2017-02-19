package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func read_json(path string, parameter interface{}) bool {
	file, err := os.Open(path)
	if err != nil {
		log.Errorln(err)
		return false
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Errorln(err)
		return false
	}
	err = json.Unmarshal(data, parameter)
	if err != nil {
		log.Errorln(err)
		return false
	}
	return true
}
