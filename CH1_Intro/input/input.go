package main

import (
	"fmt"
)

func main() {
	// Get user input
	fmt.Printf("Please give me your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("YOur name is", name)
}
