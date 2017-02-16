package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57) // Variadic function average takes variable number of arguments.
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	var total float64
	for _, v := range sf { // _ is a blank because index is not needed
		total += v
	}
	return total / float64(len(sf))
}
