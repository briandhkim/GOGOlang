package main

import "fmt"

func main() {

	var numOne int
	var numTwo int
	fmt.Print("Enter a large number: ")
	fmt.Scan(&numOne)
	fmt.Print("Enter a small number: ")
	fmt.Scan(&numTwo)

	fmt.Println("Large number divided by small number: ", numOne/numTwo)
}
