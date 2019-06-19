package main

import "fmt"

type Math interface {
	double() float64
}

func main() {
	var v Math
	vertex := Vertex{1, 2}
	v = &vertex
	fmt.Println(v.double())

	agentK := AgentK(2)
	fmt.Println(agentK.double())
	
}

type AgentK float64

func (a AgentK) double() float64 {
	return float64(a * 2)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) double() float64 {
	return (v.X + v.Y) * 2
}
