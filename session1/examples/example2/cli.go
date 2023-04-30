package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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
		fmt.Fprintln(os.Stderr, "Error: cannot open input file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: cannot read input file")
		os.Exit(1)
	}
}
