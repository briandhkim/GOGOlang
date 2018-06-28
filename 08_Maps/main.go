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
	// exampleOne()
	// compLit()
	// deleteSample()
	checkVal()
}

func exampleOne() {
	//don't do zero value creation
	//var ex map[string]int 	//this will give nil
	//with slice this is fine since slice has append, but map doesn't

	//map[key_type]value_type
	// var ex = make(map[string]int)
	// var ex = map[string]int{}
	ex := make(map[string]int)
	ex["alpha"] = 13
	ex["bravo"] = 42

	fmt.Println(ex)
}

func compLit() {
	//composite literal way of creating map
	cmp := map[string]string{
		"greeting": "hi",
		"greet":    "hello", //comma after the last map element is required for composite literal
	}
	fmt.Println(cmp)
}

func deleteSample() {
	nums := map[string]int{
		"alpha":   42,
		"bravo":   33,
		"charlie": 5,
		"delta":   8,
	}

	fmt.Println(nums)
	delete(nums, "charlie")
	fmt.Println(nums)
}

func checkVal() {
	check := map[string]int{
		"alpha":   42,
		"bravo":   33,
		"charlie": 5,
		"delta":   8,
	}

	fmt.Println(check)

	//checking if value exists
	if val, exists := check["charlie"]; exists {
		delete(check, "charlie")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("Value does not exist")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}
	fmt.Println(check)
}
