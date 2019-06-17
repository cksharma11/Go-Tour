package main

import "fmt"

func main() {
	x, y := 10, 12
	pointer := &x

	fmt.Println(*pointer)

	*pointer = 100

	fmt.Println(x)

	pointer = &y
	*pointer = *pointer * 2

	fmt.Println(y)
}
