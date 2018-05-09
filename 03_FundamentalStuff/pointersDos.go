package main

import "fmt"

func zero(x int) {
	//%p - base 16 notation with leading 0x
	fmt.Printf("address in function zero \t %p \n", &x)
	//x here is variable at different location
	x = 0
}

func memZero(x *int) {
	fmt.Printf("address in function  with pointer \t %p \n", x)
	*x = 0
}

func main() {
	x := 5
	fmt.Printf("address in main \t %p \n", &x)
	zero(x)
	fmt.Println("x here does not change ", x) //x is still 5 because 5 is only passed in as a value not the actual variable

	memZero(&x)
	fmt.Println("after using pointer, x changes", x)
}
