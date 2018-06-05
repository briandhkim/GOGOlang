package main

import "fmt"

/*
function is a type (e.g. like boolean, int, float, etc)
*/

//	callback: passing a func as an argument
func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func main() {
	//passing in an anonymous func to visit
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
}
