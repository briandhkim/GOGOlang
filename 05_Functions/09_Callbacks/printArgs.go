package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {

	var xs []int
	// xs := []int{}

	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	// fmt.Printf("%T \n", callback)
	return xs
}

//everything is passed by value in go
func main() {

	greaterThanOne := func(n int) bool {
		return n > 1
	}

	xs := filter([]int{1, 2, 3, 4}, greaterThanOne)

	//[2,3,4]
	fmt.Println(xs)
}
