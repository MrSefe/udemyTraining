package main

import "fmt"

// Each defined constant iota increases by one
const (
	A = iota
	B = iota
	C = iota * 10
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
