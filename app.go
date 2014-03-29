package main

import (
	"fmt"
)

func main() {
	var board Board
	board.Set(0, 0, &Figure{false, Rook})
	board.Set(3, 3, &Figure{false, Rook})

	board.Set(2, 2, &Figure{true, King})

	board.Print()

	for _, i := range board.Moves(2, 2) {
		fmt.Println(i)
	}
}
