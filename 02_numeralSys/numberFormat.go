package main

import "fmt"

func main() {
	fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	// %d - decimal
	// %b - binary
	// %x - hexadecimal
	fmt.Printf("%d \t %b \t %#X", 42, 42, 42)
}
