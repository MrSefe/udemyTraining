package main

import "fmt"

func main() {

	greeting := func() { // Anonymous function assigned into a variable
		fmt.Println("Hello World!")
	}

	greeting()

}
