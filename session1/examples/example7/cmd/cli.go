package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/marianina8/example7/utils"
)

func main() {
	filename := flag.String("file", "", "Input file name")
	flag.Parse()

	if *filename == "" {
		log.Fatal("Error: no input file specified")
	}

	yamlData, err := utils.Transform(*filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(yamlData))
}
