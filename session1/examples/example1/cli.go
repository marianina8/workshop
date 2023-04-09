package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	message := flag.String("message", "Hello World", "the message to print")

	flag.Parse()

	if len(flag.Args()) > 0 {
		fmt.Fprintln(os.Stderr, "Error: unexpected argument(s)")
		os.Exit(1)
	}

	fmt.Println(*message)
}
