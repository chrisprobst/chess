package main

import (
	"fmt"
	//"math/rand"
	//"time"
)

func main() {

	board, err := NewBoardFromCoords("A2", "A3", "H2", "H3", "A1", "A2", "B2", "B3")
	board, err = board.ParseAndMove("C2", "C3")
	fmt.Println(err)
	board.Print()

	/*board := NewBoard()
	board.Print()

	white := true
	for {
		moved := false
		if board.IsFoundKingChecked(white) {

			x, y, _ := board.Find(white, King)
			moves := board.Moves(x, y)

			if len(moves) == 0 {
				fmt.Println("CheckMate: ", white, " has lost!")
				break
			}

			board = board.ApplyMove(moves[rand.Intn(len(moves))])
			board.Print()
			//time.Sleep(1 * time.Second)
			moved = true
			white = !white

		} else {
		loop:
			for ny := 0; ny < 8; ny++ {
				for nx := 0; nx < 8; nx++ {
					figure := board.Get(nx, ny)
					if figure != nil && figure.White == white {
						moves := board.Moves(nx, ny)
						if len(moves) > 0 {
							board = board.ApplyMove(moves[rand.Intn(len(moves))])
							board.Print()
							//time.Sleep(1 * time.Second)
							moved = true
							white = !white

							break loop
						}
					}
				}
			}
		}
		if !moved {
			fmt.Println(white, " has lost!")
			break
		}
	}*/
}
