package api

import (
	"fmt"
)

func NewChess() Chess {
	return newBoard()
}

func NewChessFromMoves(moves ...int) (Chess, error) {
	return newBoardFromMoves(moves...)
}

func (self *board) Moves(x, y int) []*Move {
	return self.moves(x, y, false)
}

func (self *board) Move(fromX, fromY, toX, toY int) bool {
	if newBoard := self.moveOne(fromX, fromY, toX, toY); newBoard != nil {
		*self = *newBoard
		return true
	}
	return false
}

func (self *board) Print() {
	fmt.Println()
	for _, row := range self {
		for _, piece := range row {
			if piece != nil {
				piece.print()
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

type Chess interface {

	// Returns all valid moves for a given location.
	Moves(x, y int) []*Move

	// Moves the figure using the given coordinates.
	Move(fromX, fromY, toX, toY int) bool

	// Prints the board.
	Print()
}
