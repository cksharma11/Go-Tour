package main

import "fmt"

type Axis struct {
	X, Y int
}

var (
	v1 = Axis{1, 2}
	v2 = Axis{X: 0}
	v3 = Axis{}

	p = & Axis{2,4}
)

func main() {
	fmt.Println(v1,v2,v3,p)
}
