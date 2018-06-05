package main

import "fmt"

// `func() int` is a type that is being returned by `wrapper`
func wrapper() func() int {
	// x := 0
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	//func expression
	increment := wrapper()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Printf("%T \n", increment)
}


/**
closure helps limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
**/