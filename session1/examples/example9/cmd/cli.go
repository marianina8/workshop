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
	filename := flag.String("file", "", "Input file name")
	silence := flag.Bool("silence", false, "Silence the output")
	output := flag.String("output", "", "Output file name")

	flag.Parse()
	if *filename == "" {
		fmt.Fprintln(os.Stderr, "Error: no input file specified")
		os.Exit(1)
	}
	yamlData, err := utils.Transform(*filename)
	if err != nil {
		log.Fatalf("Error: unable to transform the file; unexpected format: %v", err)
		os.Exit(1)
	}
	if !*silence {
		fmt.Println(string(yamlData))
	}

	if *output != "" {
		err = ioutil.WriteFile(*output, yamlData, 0644)
		if err != nil {
			log.Fatalf("Error: unable to write to the output file: %v", err)
			os.Exit(1)
		}
	}
}
