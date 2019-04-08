package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello from go, lets run this from docker")
	if host, err := os.Hostname(); err == nil {
		fmt.Printf("Hostname: %s\n", host)
	} else {
		fmt.Printf("Error reading hostname %v", err)
	}
	fmt.Scanln()
}
