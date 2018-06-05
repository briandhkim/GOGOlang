package main

import "fmt"

//x is declared at package level scope
var x int

// var x = 0

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
}
