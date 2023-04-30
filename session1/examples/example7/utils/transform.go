package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/marianina8/example7/models"
	"gopkg.in/yaml.v2"
)

func Transform(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: cannot open input file")
		os.Exit(1)
	}
	defer file.Close()

	var jsonData models.Workshop
	err = json.NewDecoder(file).Decode(&jsonData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	yamlData, err := yaml.Marshal(&jsonData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return yamlData, nil
}
