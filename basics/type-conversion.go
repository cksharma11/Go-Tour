package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	const newConstant = "I'm a constant"
	const numericConstant = 10
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z, newConstant, numericConstant)
}
