package main

import "fmt"

func hello() {
	fmt.Println("Hello ")
}

func world() {
	fmt.Println("World!")
}

func main() {
	defer world()
	hello()
}

/*
A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body, or because the corresponding goroutine is panicking.
*/
