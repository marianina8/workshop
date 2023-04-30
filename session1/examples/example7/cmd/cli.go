package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/marianina8/example7/utils"
)

func main() {
	filename := flag.String("file", "", "Input file name")
	flag.Parse()
	if *filename == "" {
		fmt.Fprintln(os.Stderr, "Error: no input file specified")
		os.Exit(1)
	}
	yamlData, err := utils.Transform(*filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(yamlData))
}
