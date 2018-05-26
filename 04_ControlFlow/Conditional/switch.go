package main

import "fmt"

func main() {

	switch "c" {

	case "a", "b":
		fmt.Println(" case a or b")
	case "c":
		fmt.Println(" only case c")
	case "d", "e", "f":
		fmt.Println(" cases d, e, and f")
	default:
		fmt.Println("video 62 for more examples")
	}

}
