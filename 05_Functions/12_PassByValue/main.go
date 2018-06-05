package main

import "fmt"

func changeVal(z *int) {
	fmt.Println("inside change func (variable passed in as mem address)", z)
	fmt.Println("inside change func with *", *z)
	*z = 24
	fmt.Println("mem address after changing val (stays same as before)", z)
	fmt.Println("varialble at the mem address (different)", *z)
}

func main() {
	age := 55
	fmt.Println("before calling change func (mem address)", &age)

	changeVal(&age)

	fmt.Println("mem address inside main func after change func was called", &age)
	fmt.Println("value inside main func after change func executed", age)
}

/**
`age` and `z` are two different variables but point to the 
same memory address; 
the function changed a variable that was in different scope
because the value passed in was memory address and not the number value


**/