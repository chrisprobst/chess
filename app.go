package main

import (
	"fmt"
)

func main() {
	var board Board
	board.Set(0, 0, &Figure{true, Rook, Idle})
	board.Set(3, 0, &Figure{true, King, Idle})

	board.Print()

	for _, i := range board.Moves(3, 0) {
		fmt.Println(i)
	}
}
