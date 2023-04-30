package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/marianina8/example8/models"
	"gopkg.in/yaml.v2"
)

func Transform(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: cannot open input file")
		return nil, err
	}
	defer file.Close()

	var jsonData models.Workshop
	err = json.NewDecoder(file).Decode(&jsonData)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	yamlData, err := yaml.Marshal(&jsonData)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return yamlData, nil
}
