package main

import "fmt"

func main() {
	var x [256]byte

	fmt.Println(len(x)) // Just prints out the length of the array
	fmt.Println(x[42])
	for i := 0; i < 256; i++ {
		x[i] = byte(i)  // Has to bee byte
	}
	for i, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v) // Print value - type - binary. Type will actually be uint8
		if i > 50 {
			break // Stop when first 50 have been printed
		}
	}
}
