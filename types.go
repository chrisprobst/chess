package main

import (
	"errors"
	"fmt"
)

type Model int
type Transform int
type Flag int

type Figure struct {
	White bool
	Model Model
	Flags Flag
}

func (self *Figure) Print() {
	switch self.Model {
	case Pawn:
		if self.White {
			fmt.Print("P")
		} else {
			fmt.Print("p")
		}
	case Knight:
		if self.White {
			fmt.Print("N")
		} else {
			fmt.Print("n")
		}
	case Rook:
		if self.White {
			fmt.Print("R")
		} else {
			fmt.Print("r")
		}
	case Bishop:
		if self.White {
			fmt.Print("B")
		} else {
			fmt.Print("b")
		}
	case King:
		if self.White {
			fmt.Print("K")
		} else {
			fmt.Print("k")
		}
	case Queen:
		if self.White {
			fmt.Print("Q")
		} else {
			fmt.Print("q")
		}
	}
}

func NewFigure(white bool, model Model) *Figure {
	return &Figure{white, model, Idle}
}

func NewWhiteFigure(model Model) *Figure {
	return &Figure{true, model, Idle}
}

func NewBlackFigure(model Model) *Figure {
	return &Figure{false, model, Idle}
}

type Board [8][8]*Figure

func NewBoardFromCoords(coords ...string) (*Board, error) {
	return NewBoard().ParseAndMove(coords...)
}

func NewBoard() *Board {
	var board Board

	board.Set(0, 0, NewWhiteFigure(Rook))
	board.Set(1, 0, NewWhiteFigure(Knight))
	board.Set(2, 0, NewWhiteFigure(Bishop))
	board.Set(3, 0, NewWhiteFigure(King))
	board.Set(4, 0, NewWhiteFigure(Queen))
	board.Set(5, 0, NewWhiteFigure(Bishop))
	board.Set(6, 0, NewWhiteFigure(Knight))
	board.Set(7, 0, NewWhiteFigure(Rook))

	for i := 0; i < 8; i++ {
		board.Set(i, 1, NewWhiteFigure(Pawn))
	}

	board.Set(0, 7, NewBlackFigure(Rook))
	board.Set(1, 7, NewBlackFigure(Knight))
	board.Set(2, 7, NewBlackFigure(Bishop))
	board.Set(3, 7, NewBlackFigure(King))
	board.Set(4, 7, NewBlackFigure(Queen))
	board.Set(5, 7, NewBlackFigure(Bishop))
	board.Set(6, 7, NewBlackFigure(Knight))
	board.Set(7, 7, NewBlackFigure(Rook))

	for i := 0; i < 8; i++ {
		board.Set(i, 6, NewBlackFigure(Pawn))
	}

	return &board
}

const (
	Idle  = Flag(0)
	Moved = Flag(1 << 0)
)

const (
	Rook   = Model(1)
	Knight = Model(2)
	Bishop = Model(3)
	Queen  = Model(4)
	King   = Model(5)
	Pawn   = Model(6)
)

type Move struct {
	FigureOrigin, FigureDest       *Figure
	XOrigin, YOrigin, XDest, YDest int
	TransformInto                  []*Figure
	SubMove                        *Move
}

func (self *Board) Find(white bool, model Model) (int, int, bool) {
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if figure := self.Get(x, y); figure != nil &&
				figure.White == white && figure.Model == model {
				return x, y, true
			}
		}
	}
	return -1, -1, false
}

func (self *Board) ValidateWalls(x, y int) (int, int, bool) {
	if 0 <= x && 8 > x && 0 <= y && 8 > y {
		return x, y, true
	} else {
		return -1, -1, false
	}
}

func (self *Board) IsFree(x, y int) bool {
	return self.Get(x, y) == nil
}

func (self *Board) Set(x, y int, figure *Figure) bool {
	if x, y, ok := self.ValidateWalls(x, y); ok {
		(*self)[y][x] = figure
		return true
	}
	return false
}

func (self *Board) Get(x, y int) *Figure {
	if x, y, ok := self.ValidateWalls(x, y); ok {
		return (*self)[y][x]
	}
	return nil
}

func (self *Board) Print() {
	fmt.Println()
	for _, row := range self {
		for _, piece := range row {
			if piece != nil {
				piece.Print()
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (self *Board) ParseAndMove(coords ...string) (*Board, error) {
	if len(coords)%2 != 0 {
		return nil, errors.New("Number of coords must be even")
	}

	// Iterate over all coords
	for i := 0; i < len(coords); i += 2 {
		// Select strings
		from, to := coords[i], coords[i+1]

		// Extract
		fromSx, fromSy := from[:1], from[1:]
		toSx, toSy := to[:1], to[1:]

		// Parse
		fromX, fromY, err := ParseCoords(fromSx, fromSy)
		if err != nil {
			return nil, err
		}
		toX, toY, err := ParseCoords(toSx, toSy)
		if err != nil {
			return nil, err
		}

		// Move
		self = self.Move(fromX, fromY, toX, toY)
		if self == nil {
			return nil, errors.New(fmt.Sprintf("Invalid from from %s-%s to %s-%s",
				fromSx, fromSy, toSx, toSy))
		}
	}
	return self, nil
}

func (self *Board) Move(fromX, fromY, toX, toY int) *Board {

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
	return self.ApplyMove(found)
}

func (self *Board) ApplyMove(move *Move) *Board {
	// This creates a copy of the board on the stack
	copy := *self
	ptr := &copy

	for movePtr := move; movePtr != nil; movePtr = movePtr.SubMove {
		// Create copy of the struct
		copy := *movePtr.FigureOrigin
		movePtr.FigureOrigin = &copy

		if !ptr.Set(movePtr.XDest, movePtr.YDest, movePtr.FigureOrigin) {
			panic("Invalid move")
		}

		// This figure has already moved
		movePtr.FigureOrigin.Flags |= Moved

		if !ptr.Set(movePtr.XOrigin, movePtr.YOrigin, nil) {
			panic("Invalid move")
		}
	}

	return ptr
}
