package main

import "fmt"

func main() {
	// shortHandSlice()
	// varSlice()
	makeSlice()
}

func shortHandSlice() {
	student := []string{}
	students := [][]string{}
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil) //false
	//need to use append to add items
}

func varSlice() {
	var student []string
	var students [][]string
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil) //true
	//need to use append to add data
}

func makeSlice() {
	student := make([]string, 10)
	students := make([][]string, 10)
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)	//false
	//no need to use append if index is in bound of len
}

func bestWayToSlice(){
	student := make([]string, 4, 6)
	students := make([][]string, 10, 20)
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}