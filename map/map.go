package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	location := make(map[string]Vertex)
	location["Adam"] = Vertex{10,20}

	fmt.Println(location["Adam"])
}