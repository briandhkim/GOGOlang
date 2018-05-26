package main

import "fmt"

//utf8 is 4 byte coding scheme
//rune is just a character; is an alias for int32 (numberic type that stores 4 bytes)

//printing  just the value (numeric)  ;  print the value in string  ;  convert the string into byte
func main() {

	for i := 50; i <= 140; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
}
