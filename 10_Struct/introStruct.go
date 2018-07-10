package main

import "fmt"

/**
Go is Object Oriented

1) Encapsulation
	state ("fields")
	behavior ("methods")
	exported / un-exported

2) Reusability
	inheritence ("embedded types")

3) Polymorphism
	interfaces

4) Overriding
	"promotion"
**/

type person struct {
	first string	//field
	last  string	//field
	age   int 		//field
}

func main() {
	p1 := person{"John", "Smith", 32}
	// p2 := person{"Jane", "Doe", 32}

	fmt.Println(p1.first, p1.last, p1.age)
}
