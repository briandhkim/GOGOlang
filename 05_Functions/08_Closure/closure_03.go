package main

import "fmt"

func main() {
	x := 0

	//func expression; anonymous function being assigned to a variable
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())

}

/*
having narrower variable scope is better

*/
