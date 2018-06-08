package main

import "fmt"

func foo(nums ...int) {
	for _, v := range nums {
		fmt.Println(v)
	}
}

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}
