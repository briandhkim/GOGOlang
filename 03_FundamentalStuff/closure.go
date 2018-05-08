package main

import "fmt"

func main(){
	x := 42

	fmt.Println(x)
	{
		fmt.Println(x)
		y := "y's scope is inside the {}"
		fmt.Println(y)
	}

	//y is not accesible here
}