package main

import "fmt"

func main() {
	fizzBuzz(20)
}

func fizzBuzz(n int) {
	// number := n

	for i := 0; i <= n; i++ {
		if (i%5 == 0) && (i%3 == 0) {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
