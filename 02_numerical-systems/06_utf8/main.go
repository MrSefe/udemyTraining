package main

import "fmt"

func main() {
	for i := 60; i < 0; i++ {
		fmt.Printf("\t %d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
