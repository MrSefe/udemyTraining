package main

import "fmt"

func main() {
	student := []string{} // shorthand slice
	students := [][]string{} // another shorthand slice
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}