/**
there is no passing by reference in Go
everything is pass by value

slice and map are reference types
**/

package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m) // [ ]
	addName(m)
	fmt.Println(m) //[John]

	i := make(map[string]int)
	changeNum(i)
	fmt.Println(i["Jane"]) //44
}

//slice is a refernce type
//it already has a address implicitly pointing at the data
func addName(z []string) {
	z[0] = "John"
	fmt.Println(z) //[John]
}

func changeNum(z map[string]int) {
	z["Jane"] = 44
}
