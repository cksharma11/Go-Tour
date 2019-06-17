package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 1
	a[1] = 2

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	evens := [3]int{2, 4, 6}

	fmt.Println(evens)
}
