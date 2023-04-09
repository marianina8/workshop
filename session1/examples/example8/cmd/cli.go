package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/marianina8/example8/utils"
)

func main() {
	filename := flag.String("file", "", "Input file name")
	flag.Parse()

	if *filename == "" {
		log.Fatal("Error: no input file specified")
	}

	yamlData, err := utils.Transform(*filename)
	if err != nil {
		log.Fatalf("Error: unable to transform the file; unexpected format: %v", err)
		os.Exit(1)
	}

	fmt.Println(string(yamlData))
}
