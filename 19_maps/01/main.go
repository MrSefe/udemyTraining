package main

import "fmt"

func main() {
	// String as name of variable and int as type
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	// delete teh k2
	delete(m, "k2")
	fmt.Println("map:", m)

	// wtf
	v, ok := m["k2"]
	fmt.Println("ok?:", ok, v)

}
