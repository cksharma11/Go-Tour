package main

import (
	"fmt"
	"math"
)

func shortIfDemo(x float64) float64 {
	if v := math.Sqrt(9); v > x {
		return v
	} 
	return x;
}

func main() {
	fmt.Println(shortIfDemo(2))
	fmt.Println(shortIfDemo(4))
}