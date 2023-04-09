package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	filename := flag.String("file", "", "Input file name")
	flag.Parse()

	if *filename == "" {
		log.Fatal("Error: no input file specified")
	}

	yamlData, err := transform(*filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(yamlData))
}
