package main

import "fmt"

func main() {
	primes := [5]int{1, 2, 3, 5, 7}
	var firstTwoPrimes = primes[0:2]

	fmt.Println(firstTwoPrimes)
	
	firstTwoPrimes[0] = 11
	firstTwoPrimes[1] = 111
	
	fmt.Println(primes)
}
