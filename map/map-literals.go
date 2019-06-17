package main

import "fmt"

func main() {
	students := map[string]int{
		"Adam": 10,
		"Jhon": 20,
	}
	fmt.Println(students)

	elem, okk := students["Adam"]

	fmt.Println(elem, okk)

	fmt.Println(WordCount("Adam"))
}

func WordCount(word string) map[string]int {
	return map[string]int{
		word: len(word),
	}
}
