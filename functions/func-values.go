package main

import "fmt"

func addTen(a int) int {
	return a + 10
}

func addTenAndDouble(fn func(int) int, a int) int {
	value := fn(a)
	return value * 2
}

func main() {
	fmt.Println(addTenAndDouble(addTen, 2))
}