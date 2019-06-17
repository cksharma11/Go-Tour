package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v.X)
	fmt.Println(v.Y)

	v.Y = 10
	fmt.Println(v.Y)

	p := &v

	fmt.Println(p)
	(*p).X = 100
	p.Y = 200

	fmt.Println(v)
}
