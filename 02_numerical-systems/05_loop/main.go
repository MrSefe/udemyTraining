package main

import "fmt"

func main() {
	for i := 0; i < 20000; i++ {
		fmt.Printf("\t %d \t %b \t %x \n", i, i, i)
	}
}
