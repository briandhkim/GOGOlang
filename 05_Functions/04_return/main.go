package main

import "fmt"

func main() {
	s := greet("Jane", "Doe")
	fmt.Println(s)
}

//Sprint = string print
//printing to a string (i.e. instead of printing out to monitor, basically concatenates the string)
func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
