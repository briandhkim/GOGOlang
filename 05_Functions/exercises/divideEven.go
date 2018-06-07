package main

import "fmt"

func main() {
	hFour, hFourEven := half(4)
	hSeven, hSevenEven := half(7)

	fmt.Println("input of 4: ", hFour, hFourEven)
	fmt.Println("input of 7: ", hSeven, hSevenEven)
}

func half(num int) (int, bool) {
	x := num / 2
	return x, x%2 == 0
}
