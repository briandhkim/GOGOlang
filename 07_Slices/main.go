package main

import "fmt"

/**

**/

func main() {
	slicingExample()
}

func slicingExample() {
	strSlice := []string{"a", "b", "c", "d", "e", "f"}
	//slicing
	fmt.Println("Slicing example: ", strSlice[2:4])    //will give items at index 2 and 3
	fmt.Println("index access example: ", strSlice[2]) //index access
	fmt.Println("Index access of string"[2])           //index access with string
	/*
		'slicing' on string can be done because a string is a slice of bytes
		- a string is made up of runes
		- a rune is a unicode code point
		- a unicode code point is 1 to 4 bytes
		strings are made up of runes, runes are made up of bytes, so strings are made up
		of bytes. a string is a bunch of bytes; a slice of bytes
	*/
}

