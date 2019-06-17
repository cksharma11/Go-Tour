package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	var v = a[0:5]
	fmt.Println(v)
	v = a[0:]
	fmt.Println(v)
	v = a[:5]
	fmt.Println(v)
	v = a[:]
	fmt.Println(v)

	v = a[0:2]
	fmt.Println(len(v))
	fmt.Println(cap(v))
}
