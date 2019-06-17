package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) abs() float64 {
	return math.Sqrt(v.X + v.Y)
}

func main() {
	var vertex = Vertex{5,4}
	fmt.Println(vertex.abs())
}
