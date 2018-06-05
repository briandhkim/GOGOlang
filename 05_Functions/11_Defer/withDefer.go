package main

import "fmt"

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}


func main() {
	//defer the execution of `world` until right
	//before the `main` function exits
	//often used to close a file that was opened
	defer world()
	hello()
}
