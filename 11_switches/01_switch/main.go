package main

import "fmt"

func main() {

	switch "Medhi" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medhi":
		fmt.Println("Wassup Medhi")
		fallthrough //drops to the next case
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default: // this is the default fallthrough
		fmt.Println("No names matched!")
	}
}

/*
	There's no default fallthrough. It must be defined if needed.
*/
