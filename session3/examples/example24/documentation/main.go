package main

import (
	"log"

	"github.com/marianina8/example24/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	header := &doc.GenManHeader{
		Title:  "JSON to YAML converter",
		Source: "Auto generated by marianina8",
	}
	err := doc.GenManTree(cmd.RootCmd(), header, "./docs")
	if err != nil {
		log.Fatal(err)
	}
}
