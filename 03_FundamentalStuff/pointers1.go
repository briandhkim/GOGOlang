package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	//b is of type 'int pointer'
	var b *int = &a

	fmt.Println("b is referencing memory address ", b)
	fmt.Println("*b gives memory reference value(dereference) ", *b)

}
