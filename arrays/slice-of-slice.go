package main

import (
	"fmt"
	"strings"
)

func main() {
	//create a ttt board
	board := [][]string {
		[]string{"_","_","_"},
		[]string{"_","_","_"},
		[]string{"_","_","_"},
	}

	//The player take turn

	board[0][0] = "X"
	board[0][2] = "O"
	board[1][1] = "X"
	board[2][1] = "O"
	board[0][1] = "O"

	for i := 0; i < len(board); i++ {
		fmt.Println("%s\n", strings.Join(board[i], " "))
	}
}