package main

import "fmt"

func main() {

	mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Wednesday", "Thursday", "Friday"}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)

	// Moving slice to undefined place deletes the slice.
	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println(mySlice)
}
