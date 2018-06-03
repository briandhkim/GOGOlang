package main

import "fmt"

func main() {

	//initialization
	//keeps the scope of food variable inside the if statement
	b := true
	if food := "Pie"; b {
		fmt.Println(food)
	}

	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

}
