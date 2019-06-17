package main

import "fmt"

func createAdd(a int) func(int) int {
	return func(b int) int {
		return b + a
	}
}

func main() {
	var addTen = createAdd(10)
	var v = addTen(10)

	var addTwenty = createAdd(20)
	var x = addTwenty(10)

	fmt.Println(v)
	fmt.Println(x)
}