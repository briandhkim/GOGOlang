package main

//`fmt` is at a file level
import "fmt"

/**
thinking about scope: what outer scope is enclosing inner scope
this is an example looking at package level scope
**/
func main() {
	//block level scope
	x := 42
	fmt.Println("x is at the outer scope", x)
	{
		fmt.Println("x is accessible here", x)
		y := "y is in the inner scope"
		fmt.Println(y)
	}
	// fmt.Println(y) //outside scope of y

}
