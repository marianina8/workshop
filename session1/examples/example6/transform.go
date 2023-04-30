package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func transform(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: cannot open input file")
		os.Exit(1)
	}
	defer file.Close()

	var jsonData Workshop
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
