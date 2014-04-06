package api

import (
	"errors"
	"fmt"
)

var (
	invalidNumberOfCoords error = errors.New(
		"Number of coords must be divisible by 4")
)

type Model int
type transform int
type flag int

type Figure struct {
	White bool
	Model Model

	flags flag
}

func (self *Figure) print() {
	switch self.Model {
	case pawn:
		if self.White {
			fmt.Print("P")
		} else {
			fmt.Print("p")
		}
	case knight:
		if self.White {
			fmt.Print("N")
		} else {
			fmt.Print("n")
		}
	case rook:
		if self.White {
			fmt.Print("R")
		} else {
			fmt.Print("r")
		}
	case bishop:
		if self.White {
			fmt.Print("B")
		} else {
			fmt.Print("b")
		}
	case king:
		if self.White {
			fmt.Print("K")
		} else {
			fmt.Print("k")
		}
	case queen:
		if self.White {
			fmt.Print("Q")
		} else {
			fmt.Print("q")
		}
	}
}

func newFigure(white bool, model Model) *Figure {
	return &Figure{white, model, idle}
}

func newWhiteFigure(model Model) *Figure {
	return &Figure{true, model, idle}
}

func newBlackFigure(model Model) *Figure {
	return &Figure{false, model, idle}
}

type board [8][8]*Figure

func newBoard() *board {
	var board board

	board.set(0, 0, newWhiteFigure(rook))
	board.set(1, 0, newWhiteFigure(knight))
	board.set(2, 0, newWhiteFigure(bishop))
	board.set(3, 0, newWhiteFigure(king))
	board.set(4, 0, newWhiteFigure(queen))
	board.set(5, 0, newWhiteFigure(bishop))
	board.set(6, 0, newWhiteFigure(knight))
	board.set(7, 0, newWhiteFigure(rook))

	for i := 0; i < 8; i++ {
		board.set(i, 1, newWhiteFigure(pawn))
	}

	board.set(0, 7, newBlackFigure(rook))
	board.set(1, 7, newBlackFigure(knight))
	board.set(2, 7, newBlackFigure(bishop))
	board.set(3, 7, newBlackFigure(king))
	board.set(4, 7, newBlackFigure(queen))
	board.set(5, 7, newBlackFigure(bishop))
	board.set(6, 7, newBlackFigure(knight))
	board.set(7, 7, newBlackFigure(rook))

	for i := 0; i < 8; i++ {
		board.set(i, 6, newBlackFigure(pawn))
	}

	return &board
}

func newBoardFromMoves(coords ...int) (*board, error) {
	return newBoard().moveMultiple(coords...)
}

const (
	idle  = flag(0)
	moved = flag(1 << 0)
)

const (
	rook   = Model(1)
	knight = Model(2)
	bishop = Model(3)
	queen  = Model(4)
	king   = Model(5)
	pawn   = Model(6)
)

type Move struct {
	FigureOrigin, FigureDest       *Figure
	XOrigin, YOrigin, XDest, YDest int
	TransformInto                  []*Figure
	SubMove                        *Move
}

func (self *board) find(white bool, model Model) (int, int, bool) {
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if figure := self.get(x, y); figure != nil &&
				figure.White == white && figure.Model == model {
				return x, y, true
			}
		}
	}
	return -1, -1, false
}

func (self *board) validateWalls(x, y int) (int, int, bool) {
	if 0 <= x && 8 > x && 0 <= y && 8 > y {
		return x, y, true
	} else {
		return -1, -1, false
	}
}

func (self *board) isFree(x, y int) bool {
	return self.get(x, y) == nil
}

func (self *board) set(x, y int, figure *Figure) bool {
	if x, y, ok := self.validateWalls(x, y); ok {
		(*self)[y][x] = figure
		return true
	}
	return false
}

func (self *board) get(x, y int) *Figure {
	if x, y, ok := self.validateWalls(x, y); ok {
		return (*self)[y][x]
	}
	return nil
}

func (self *board) moveMultiple(coords ...int) (*board, error) {
	if len(coords)%4 != 0 {
		return nil, invalidNumberOfCoords
	}

	// Iterate over all coords
	for i := 0; i+3 < len(coords); i += 4 {
		// Select coordinates
		fromX, fromY := coords[i], coords[i+1]
		toX, toY := coords[i+2], coords[i+3]

		// Move
		self = self.moveOne(fromX, fromY, toX, toY)
		if self == nil {
			return nil, errors.New(fmt.Sprintf("Invalid from from %d-%d to %d-%d",
				fromX, fromY, toX, toY))
		}
	}
	return self, nil
}

func (self *board) moveOne(fromX, fromY, toX, toY int) *board {

	// Look through all valid moves and select valid
	var found *Move
	for _, move := range self.Moves(fromX, fromY) {
		if move.XDest == toX && move.YDest == toY {
			found = move
			break
		}
	}

	// There is no such much
	if found == nil {
		return nil
	}

	// Apply the found move
	return self.applyMove(found)
}

func (self *board) applyMove(move *Move) *board {
	// This creates a copy of the board on the stack
	copy := *self
	ptr := &copy

	for movePtr := move; movePtr != nil; movePtr = movePtr.SubMove {
		// Create copy of the struct
		copy := *movePtr.FigureOrigin
		movePtr.FigureOrigin = &copy

		if !ptr.set(movePtr.XDest, movePtr.YDest, movePtr.FigureOrigin) {
			panic("Invalid move")
		}

		// This figure has already moved
		movePtr.FigureOrigin.flags |= moved

		if !ptr.set(movePtr.XOrigin, movePtr.YOrigin, nil) {
			panic("Invalid move")
		}
	}

	return ptr
}
