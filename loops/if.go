package main

import "fmt"

func getLargerNumber(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	fmt.Println(getLargerNumber(10, 11))
}
