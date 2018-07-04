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
	// checkVal()
	rangeLoop()
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

//checking if value exists
func checkVal() {
	check := map[string]int{
		"alpha":   42,
		"bravo":   33,
		"charlie": 5,
		"delta":   8,
	}

	fmt.Println(check)

	// delete(check, "charlie")

	// 'comma ok' idiom;  value, ok := map['key']
	//value = map[key] value, ok = bool
	//val, exists := check["charlie"];	//instead of keeping this in separate line, having it in the `if` line keeps the scope confined
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

	// value, ex := check["bravo"]
	// fmt.Printf("%T %T \n", ex, value)
}

func rangeLoop() {
	names := map[int]string{
		0: "henry",
		1: "bob",
		2: "jane",
		3: "sam",
	}

	for key, val := range names {
		fmt.Println(key, "--", val)
	}
}
