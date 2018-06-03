package main

import "fmt"

func main() {
	var userNumber int
	fmt.Print("Enter a number between 1 and 100: ")
	fmt.Scan(&userNumber)

	fizzBuzz(userNumber)
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
