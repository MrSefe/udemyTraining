package main

import "fmt"

func main() {

	var numOne int
	var numTwo int
	fmt.Println("Please enter a number: ")
	fmt.Scan(&numOne)
	fmt.Println("Please enter a number: ")
	fmt.Scan(&numTwo)
	fmt.Println(numOne, "/", numTwo, "=", numOne / numTwo)
}