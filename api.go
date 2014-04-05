package main

import (
	"fmt"
)

func NewChess() Chess {
	return newBoard()
}

func NewChessFromMoves(moves ...int) (Chess, error) {
	return newBoardFromMoves(moves...)
}

type Move struct {
	XOrigin, YOrigin, XDest, YDest int
	TransformInto                  []int
	SubMove                        *Move
}

func (self *move) toMove() *Move {
	if self == nil {
		return nil
	}

	newMove := &Move{
		self.xOrigin, self.yOrigin, self.xDest, self.yDest,
		make([]int, len(self.transformInto)),
		self.subMove.toMove(),
	}

	for idx, val := range self.transformInto {
		newMove.TransformInto[idx] = int(val.model)
	}

	return newMove
}

func (self *board) ComputeMoves(x, y int) []*Move {
	moves := self.moves(x, y)
	newMoves := make([]*Move, len(moves))
	for idx, val := range moves {
		newMoves[idx] = val.toMove()
	}
	return newMoves
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

	/*
		Returns all valid moves for a given location.
	*/
	ComputeMoves(x, y int) []*Move

	/*
		Moves the figure using the given coordinates.
	*/
	Move(fromX, fromY, toX, toY int) bool

	/*
		Prints the board.
	*/
	Print()
}
