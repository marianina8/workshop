package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := flag.String("file", "", "Input file name")
	flag.Parse()
	if *filename == "" {
		fmt.Fprintln(os.Stderr, "Error: no input file specified")
		os.Exit(1)
	}
	yamlData, err := transform(*filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(yamlData))
}
