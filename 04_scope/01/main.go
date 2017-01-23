package main

import "fmt"

func main() {
	x := 0
	// Anonymous function = function with out a name
	increment := func() int { // Function expression = assigning function to a variable
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
