package main

import "fmt"

func createFiboGenerator() func(a int) int {
	return func(a int) int {
		if a == 0 || a == 1 {
			return a
		}
		//Yehh it's a bad one but for this purpose now ok
		return createFiboGenerator()(a - 1) + createFiboGenerator()(a - 2)
	}
}

func main() {
	var fiboGenerator = createFiboGenerator();
	fmt.Println(fiboGenerator(5))
}