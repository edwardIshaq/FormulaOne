package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Yo yo yo")
	if host, err := os.Hostname(); err == nil {
		fmt.Printf("Hostname: %s\n", host)
	} else {
		fmt.Printf("Error reading hostname %v", err)
	}
	fmt.Scanln()
}
