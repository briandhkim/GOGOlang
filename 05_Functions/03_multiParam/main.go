package main

import "fmt"

func main() {
	lastName := "Smith"
	fullName("John", lastName)
}

func fullName(fname string, lname string) {
	fmt.Println("first name: ", fname, ",  last name: ", lname)
}
