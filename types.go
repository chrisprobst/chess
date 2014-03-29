package main

import (
	"fmt"
)

type Model int
type Transform int

type Figure struct {
	white bool
	model Model
}

type Board [8][8]*Figure

const (
	Rook   = Model(1)
	Knight = Model(2)
	Bishop = Model(3)
	Queen  = Model(4)
	King   = Model(5)
	Pawn   = Model(6)
)

const (
	NoTransform       = Transform(0)
	PawnTransform     = Transform(1)
	KingRookTransform = Transform(2)
)

type Move struct {
	FigureOrigin, FigureDest       *Figure
	XOrigin, YOrigin, XDest, YDest int
	TransformMove                  Transform
}

func (self *Board) Find(white bool, model Model) (int, int, bool) {
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if figure := self.Get(x, y); figure != nil &&
				figure.white == white && figure.model == model {
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
				fmt.Print(piece.model)
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (self *Board) ApplyMove(move *Move) *Board {
	// This creates a copy of the board on the stack
	copy := *self
	ptr := &copy

	if !ptr.Set(move.XDest, move.YDest, move.FigureOrigin) {
		panic("Invalid move")
	}

	if !ptr.Set(move.XOrigin, move.YOrigin, nil) {
		panic("Invalid move")
	}

	return ptr
}
