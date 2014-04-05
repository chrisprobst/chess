package main

import (
	"encoding/json"
	"fmt"
	//"math/rand"
	//"time"
)

func main() {

	chess, _ := NewChessFromMoves(
		0, 1, 0, 2,
	)

	chess.Print()

	moves := chess.ComputeMoves(2, 1)
	b, _ := json.Marshal(moves)
	fmt.Println(string(b))

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
