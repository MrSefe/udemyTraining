package main

import (
	"fmt"
)

func main() {

	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Hyvää päivää!",
	}

	// Looping through using range
	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
}
