package main

import "fmt"

func main() {
	// This may be considered as the best way (used in lang spec)
	student := make([]string, 35)
	students := make([][]string, 35)
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}
