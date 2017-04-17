package main

import "fmt"

func main() {
	fmt.Println(findHighest(5, 6, 24, 3, 8, 13, 0, -2))
}

// Takes series of numbers and returns the highest one
func findHighest(numbers ...int) int {

	var highest int
	for _, v := range numbers {
		if v > highest {
			highest = v
		}

	}
	return highest

}
