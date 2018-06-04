package main
//video 71 Func Expressions
import "fmt"

func main() {

	//only way to have a function inside another
	//can't do `func greeting(){ ... }`
	greeting := func() {
		fmt.Println("hello world")
	}

	greeting()
	// fmt.Printf("%T \n", greeting)

	greetingTwo := anotherExpression()
	fmt.Println(greetingTwo())
	// fmt.Printf("%T \n", greetingTwo)
}

func anotherExpression() func() string {
	//returning an anonymous function that returns a string
	return func() string {
		return "Another way of doing function expression"
	}
}
