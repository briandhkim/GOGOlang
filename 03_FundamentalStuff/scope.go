package main

import "fmt"

//package level scope
var x int = 42 

func main(){
	fmt.Println(x)

	//block level scope
	y := 7
	fmt.Println(y)

	foo()
}

func foo(){
	fmt.Println(x)

	// y is not accesible here
}