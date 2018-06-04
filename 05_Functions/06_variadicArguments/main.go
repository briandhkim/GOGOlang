package main

import "fmt"

func main() {
	//this is slice; a data type like an array
	data := []float64{43, 56, 87, 12, 45, 57}

	n := average(data...) // `...` opens up the `data` slice to individual float64 variables
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	//fmt.Println(sf)
	//fmt.Printf("%T \n", sf)

	var total float64 //set to 0 value (i.e. 0.0)

	//for index,value := ...
	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}
