package main

import (
	"fmt"
	"math"
)

type Coords struct {
	X, Y float64
}

func (c Coords) Abs() float64{
	return math.Sqrt(c.X + c.Y)
}

func (c *Coords) Scale(f float64) {
	c.X = c.X * f
	c.Y = c.Y * f
}

func main() {
	c := Coords{1,2}
	c.Scale(10)
	fmt.Println(c)
	fmt.Println(c.Abs())
}