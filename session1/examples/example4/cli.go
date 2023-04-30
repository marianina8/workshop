package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Instructor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Workshop struct {
	Title       string       `json:"title"`
	Instructors []Instructor `json:"instructors"`
}

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

	var jsonData Workshop
	err = json.NewDecoder(file).Decode(&jsonData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Workshop: %s\n", jsonData.Title)
	fmt.Println("Instructors:")
	for _, instructor := range jsonData.Instructors {
		fmt.Printf("  %s (%s)\n", instructor.Name, instructor.Email)
	}
}
