package rules

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
	fmt.Print(self.SPrint())
}

func (self *board) SPrint() string {
	s := fmt.Sprintln()
	for _, row := range self {
		for _, piece := range row {
			if piece != nil {
				s += piece.sprint()
			} else {
				s += fmt.Sprint(".")
			}
		}
		s += fmt.Sprintln()
	}
	s += fmt.Sprintln()
	return s
}

func (self *board) Board() *[8][8]Model {
	grid := new([8][8]Model)
	for y, row := range *self {
		for x, column := range row {
			if column != nil {
				grid[y][x] = column.Model
			}
		}
	}
	return grid
}

type Chess interface {

	// Returns all valid moves for a given location.
	Moves(x, y int) []*Move

	// Moves the figure using the given coordinates.
	Move(fromX, fromY, toX, toY int) bool

	// Return the board for rendering
	Board() *[8][8]Model

	// DEBUG: Prints the board.
	Print()
	SPrint() string
}
