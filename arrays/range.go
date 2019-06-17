package main

import "fmt"

func main() {
	pow := []int{1, 2, 3, 4, 5}

	//using index and value
	for i, v := range pow {
		fmt.Println(i+1, v*v)
	}

	//Skip index using _
	for _, v := range pow {
		fmt.Println(v)
	}

	//just use index
	for i := range pow {
		fmt.Println(i)
	}
}
