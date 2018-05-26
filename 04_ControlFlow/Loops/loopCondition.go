package main

import "fmt"

func main() {
	i := 0

	//same as 'while' but go doens't use 'while'
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for {
	// 	fmt.Println(i)
	// 	if i >= 10 {
	// 		break
	// 	}
	// 	i++
	// }

	// only printing odd numbers
	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}
}
