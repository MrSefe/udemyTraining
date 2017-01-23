package main

import "fmt"

func main() {
	x := 5
	zero(x)        // passes the value to function, not the variable
	fmt.Println(x) // x is still 5
	two(&x)        // using & to pass the address (pointer)
	fmt.Println(x) // now the x has changed
}

func zero(x int) { // this function get only the value
	x = 0 // this is not the same variable x
}

func two(x *int) { // this get the actual pointer (address) of x
	*x = 2 // this points to the x and so it changes
}
