package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/marianina8/example9/utils"
)

func main() {
	filename := flag.String("input", "", "Input file name")
	silence := flag.Bool("silence", false, "Silence output")
	output := flag.String("output", "", "Input file name")

	flag.Parse()

	if *filename == "" {
		log.Fatal("Error: no input file specified")
	}

	yamlData, err := utils.Transform(*filename)
	if err != nil {
		log.Fatalf("Error: unable to transform the file; unexpected format: %v", err)
		os.Exit(1)
	}

	if *silence {
		os.Exit(0)
	}

	if *output == "" {
		fmt.Println(string(yamlData))
	} else {
		err = ioutil.WriteFile(*output, yamlData, 0644)
		if err != nil {
			log.Fatalf("Error: unable to write output file: %v", err)
			os.Exit(1)
		}
	}

}
