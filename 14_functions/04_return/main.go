package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe")) // Function greet return a string
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

// Sprint print into a string