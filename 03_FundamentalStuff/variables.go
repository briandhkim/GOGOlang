package main

import "fmt"

func main(){
	//shorthand variable declare/assign/initialize
	a := 10
	b := "go"
	c := 3.5
	d := true

	//%v -  value in defualt format
	//%T - value type
	fmt.Printf("%v %T \n", a, a)
	fmt.Printf("%v %T \n", b, b)
	fmt.Printf("%v %T \n", c, c)
	fmt.Printf("%v %T \n", d, d)

	//declare to zero value
	var e int
	var f string
	var g float64
	var h bool

	fmt.Printf("%v %T \n", e, e)
	fmt.Printf("%v %T \n", f, f)
	fmt.Printf("%v %T \n", g, g)
	fmt.Printf("%v %T \n", h, h)

}