package main

type Figure struct {
	white bool
	model int
}

type Board [8][8]*Figure

const (
	Rook   = 1
	Knight = 2
	Bishop = 3
	Queen  = 4
	King   = 5
	Pawn   = 6
)

type Move struct {
	FigureOrigin, FigureDest       *Figure
	XOrigin, YOrigin, XDest, YDest int
	TransformMove                  bool
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
