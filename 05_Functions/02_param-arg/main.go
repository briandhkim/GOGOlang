package main

import "fmt"

func main() {
	greet("John")
	greet("Jane")
}

//no receiver or return
func greet(name string) {
	fmt.Println(name)
}

//greet is declared with a ****parameter****
//when calling greet, pass in an ****argument****
