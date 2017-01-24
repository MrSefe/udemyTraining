package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue // this skips rest of the loop but _continues_ looping
		}
		fmt.Println(i)
		if i >= 50 {
			break // this breaks te loop
		}
	}
}
