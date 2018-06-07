package main

import "fmt"

func main() {
	half := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}

	hFour, hFourEven := half(4)
	hSeven, hSevenEven := half(7)

	fmt.Println("input of 4: ", hFour, hFourEven)
	fmt.Println("input of 7: ", hSeven, hSevenEven)

}
