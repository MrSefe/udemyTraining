package main

import "fmt"

// Constants are NOT capitalized! You can use co_ to mark them up.
const p string = "death & taxes"

func main() {

	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
}

// a CONSTANT is a simple unchanging value
