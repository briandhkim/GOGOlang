package main

import "fmt"

/**
maps:
- key/value storage
- a 'dictionary'
- an unordered group of elements of one type, called the element type
	- indexed by a set of unique keys of another type, called the key type
	- vlue of an uninitialized map is nil (because map is reference type like slice)

a map is built on top of a hash table
map is reference type
	stores a pointer to the address where the data is actually stored

**/

func main() {
	exampleOne()
}

func exampleOne() {
	//don't do zero value creation
	//var ex map[string]int 	//this will give nil

	//map[key_type]value_type
	// var ex = make(map[string]int)
	// var ex = map[string]int{}
	ex := make(map[string]int)
	ex["alpha"] = 13
	ex["bravo"] = 42

	fmt.Println(ex)
}
