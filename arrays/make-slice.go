package main

import "fmt"

func main() {
	a := make([]int, 5)
	fmt.Println(a)

	b := make([]int, 1, 5)
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))
}