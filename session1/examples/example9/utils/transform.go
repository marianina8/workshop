package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/marianina8/example9/models"
	"gopkg.in/yaml.v2"
)

func Transform(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var jsonData models.Workshop
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
