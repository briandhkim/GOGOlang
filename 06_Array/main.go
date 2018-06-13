package main

import "fmt"

/*

array:
var x [58]string

slice:
var x []string

*/

func main() {
	// arrSample_1()
	arrSample_2()
}

func arrSample_1() {
	var x [58]string
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
}

func arrSample_2() {
	var x [100]int
	fmt.Println(len(x))
	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 0; i < 100; i++ {
		x[i] = i
	}
	for i, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 50 {
			break
		}
	}
}
