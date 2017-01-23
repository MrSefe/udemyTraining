package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a


	fmt.Println(b) // references to b's memory address and gives the address
	fmt.Println(*b) // dereferences to location and gives the value

	/*

		The above code makes b a pointer to the memory address where variable a is stored. B is of type "int pointer".
		*int -- the * is part of the type -- bis of type *int.

	*/
}
