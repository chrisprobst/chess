package main

import (
	"fmt"
)

type Model int
type Transform int
type Flag int

type Figure struct {
	white bool
	model Model
	flags Flag
}

type Board [8][8]*Figure

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

	for movePtr := move; movePtr != nil; movePtr = movePtr.SubMove {
		// Create copy of the struct
		copy := *movePtr.FigureOrigin
		movePtr.FigureOrigin = &copy

		if !ptr.Set(movePtr.XDest, movePtr.YDest, movePtr.FigureOrigin) {
			panic("Invalid move")
		}

		// This figure has already moved
		movePtr.FigureOrigin.flags |= Moved

		if !ptr.Set(movePtr.XOrigin, movePtr.YOrigin, nil) {
			panic("Invalid move")
		}
	}

	return ptr
}
