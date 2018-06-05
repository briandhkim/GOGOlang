package main

import "fmt"

func main() {

	func() {
		fmt.Println("Executing by itself because () added")
	}()
	//added () at the end for execution
}
