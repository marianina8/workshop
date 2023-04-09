package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	filename := flag.String("file", "", "Input file name")

	flag.Parse()

	if *filename == "" {
		fmt.Fprintln(os.Stderr, "Error: no input file specified")
		os.Exit(1)
	}

	file, err := os.Open(*filename)
	if err != nil {
		fmt.Println(err)
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
	fmt.Println(string(yamlData))

}
