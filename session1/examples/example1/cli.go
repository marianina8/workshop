package main

import (
	"flag"
	"fmt"
)

func main() {
	message := flag.String("message", "Hello World", "Message to print")
	flag.Parse()
	if (len(flag.Args()) > 0) {
		fmt.Println("Invalid arguments")
		return
	}
	fmt.Println(*message)
}
