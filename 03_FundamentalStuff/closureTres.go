package main

import "fmt"

func main() {
	//no longer a package level scope
	x := 0

	//anonymous function
	//func expression - assigning function to a variable
	//can't create a function inside another function? check video 44
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
}
