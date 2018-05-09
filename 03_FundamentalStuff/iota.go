package main

import "fmt"

const (
	A = iota //0
	B        //1
	C        //2
)

const (
	D = iota //0
	E        //1
	F        //2
)

const (
	_	=	iota	//0
	KB	=	1 << (iota * 10) //1 << (1*10)
	MB	=	1 << (iota * 10) //1 << (2*10)
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)

	fmt.Printf("binary\t\tdecimal\n")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
}
