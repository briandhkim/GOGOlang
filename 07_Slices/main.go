package main

import "fmt"

/**
	slice has built-in function `len` for getting length
	length of slice can change, unlike arrays

	capacity: the array underlying a slice may extend past
		the end of the slice. capacity is a measure of that
		extent; it is the sum of the length of the slice and
		the length of the array beyond the slice; a slice of
		length up to that capacity can be created by slicing a
		new one from the original slice. The capacity of a
		slice can be discovered using the built-in function `cap`

	slice can be created with `make` function:
		`make( []T, length, capacity )`		//T is likely the type property (e.g. string, int, etc)
	or using `new`
		`new( [100]int )[0:50]`
**/

func main() {
	slicingExample()
}

func slicingExample() {
	strSlice := []string{"a", "b", "c", "d", "e", "f"}
	//slicing
	fmt.Println("Slicing example: ", strSlice[2:4]) //will give items at index 2 and 3
	//starts reading at 2 and stops at 4(not including)
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
