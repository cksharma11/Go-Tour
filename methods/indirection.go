package main

import "fmt"

type Coordinates struct {
	X, Y float64
}

func (c Coordinates) Scale(f float64) {
	c.X = c.X * f
	c.Y = c.Y * f
}

func ScaleFunc(c *Coordinates, f float64) {
	c.X = c.X * f
	c.Y = c.Y * f
}

func main() {
	c := Coordinates{1, 2}
	c.Scale(10)

	fmt.Println(c)

	p := &c

	ScaleFunc(p, 10)
	fmt.Println(p)
}
