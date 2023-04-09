package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func transform(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var jsonData Workshop
	err = json.Unmarshal(bytes, &jsonData)
	if err != nil {
		return nil, err
	}

	yamlData, err := yaml.Marshal(&jsonData)
	if err != nil {
		return nil, err
	}

	return yamlData, nil
}
