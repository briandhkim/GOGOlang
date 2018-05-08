package main

import "fmt"

//wrapper() -function name and parameter
//func() int -function returns a function
// i.e. returns 'function type'
func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()

	fmt.Println(increment())
	fmt.Println(increment())
}
