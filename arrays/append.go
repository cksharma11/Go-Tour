package main

import "fmt"

func printSlice(s []int) {
	fmt.Println("len=%d cap=%d %v \n", len(s), cap(s), s)
}

func main() {
	var s []int
	printSlice(s)

	//append works on nil slices
	s = append(s, 1)
	printSlice(s)

	//can add more than one elements at a time
	s = append(s, 1,2,3,4)
	printSlice(s)
}
