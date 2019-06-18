package main

import "fmt"

func hasWon(moves []int) bool {
	winningMoves := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	for i := 0; i < len(winningMoves); i++ {
		result := true
		for j := 0; j < len(winningMoves); j++ {
			result = isIncludes(moves, winningMoves[i][j])
			if result == false {
				break
			}
		}
		if result == true {
			return true
		}
	}
	return false
}

func isIncludes(list []int, x int) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == x {
			return true
		}
	}
	return false
}

func main() {
	moves := []int{1, 2, 3}
	cMoves := []int{2, 3, 5}
	dMoves := []int{2, 3, 5, 1}

	fmt.Println("true ==", hasWon(moves))
	fmt.Println("false ==", hasWon(cMoves))
	fmt.Println("true ==", hasWon(dMoves))
}
