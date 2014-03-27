package main

import (
	"fmt"
)

func (self *Board) PawnMoves(x, y int) (moves []*Move) {

	// Get the figure
	figure := self.Get(x, y)
	if figure == nil || figure.model != Pawn {
		return
	}

	// Color to direction
	direction := 1
	if !figure.white {
		direction = -1
	}

	// The pawn is at the end of the trail
	if _, _, ok := self.ValidateWalls(x, y+direction); !ok {
		return
	}

	// Check if this pawn would transform into something
	transformMove := false
	if newLoc := y + direction; newLoc == 0 || newLoc == 7 {
		transformMove = true
	}

	// Check usual forward movement
	if self.IsFree(x, y+direction) {

		// Append the forward move
		moves = append(moves, &Move{figure, nil, x, y, x, y + direction, transformMove})

		// Check double forward movement
		if (y == 1 || y == 6) && self.IsFree(x, y+2*direction) {

			// Append double move
			moves = append(moves, &Move{figure, nil, x, y, x, y + 2*direction, false})
		}
	}

	// Check if we can hit an enemy
	for _, offset := range []int{-1, 1} {
		if enemy := self.Get(x+offset, y+direction); enemy != nil && enemy.white != figure.white {
			moves = append(moves, &Move{figure, enemy, x, y, x + offset, y + direction, transformMove})
		}
	}

	return
}

func (self *Board) KnightMoves(x, y int) (moves []*Move) {
	// Get the figure
	figure := self.Get(x, y)
	if figure == nil || figure.model != Knight {
		return
	}

	for _, direction := range [][]int{
		[]int{0, 1},
		[]int{1, 0},
		[]int{0, -1},
		[]int{-1, 0},
	} {
		for _, variation := range []int{-1, 1} {
			// Calculate base heading
			nx, ny := x+2*direction[0], y+2*direction[1]

			// Add variance
			nx, ny = nx+direction[1]*variation, ny+direction[0]*variation

			// Check if the position is inside the board
			if _, _, ok := self.ValidateWalls(nx, ny); ok {

				// Get the figure at this location
				enemy := self.Get(nx, ny)

				// If there is no figure or the figure is an enemy
				if enemy == nil || enemy.white != figure.white {

					// Append this move
					moves = append(moves, &Move{figure, enemy, x, y, nx, ny, false})
				}
			}
		}
	}
	return

	return nil
}

func (self *Board) RookMoves(x, y int) (moves []*Move) {
	// Get the figure
	figure := self.Get(x, y)
	if figure == nil || figure.model != Rook {
		return
	}

	// Going vertically
	for _, direction := range []int{-1, 1} {

		// Walk into direction
		for ny := y + direction; ny >= 0 && ny < 8; ny += direction {

			// Check board bounds
			if _, _, ok := self.ValidateWalls(x, ny); ok {
				// Get figure
				enemy := self.Get(x, ny)

				// If there is no figure or the figure is an enemy
				if enemy == nil || enemy.white != figure.white {
					moves = append(moves, &Move{figure, enemy, x, y, x, ny, false})
				}

				if enemy != nil {
					break
				}
			} else {
				break
			}
		}
	}

	// Going horizontally
	for _, direction := range []int{-1, 1} {
		// Walk into direction
		for nx := x + direction; nx >= 0 && nx < 8; nx += direction {
			// Check board bounds
			if _, _, ok := self.ValidateWalls(nx, y); ok {
				// Get figure
				enemy := self.Get(nx, y)

				// If there is no figure or the figure is an enemy
				if enemy == nil || enemy.white != figure.white {
					moves = append(moves, &Move{figure, enemy, x, y, nx, y, false})
				}

				if enemy != nil {
					break
				}
			} else {
				break
			}
		}
	}

	return
}

func (self *Board) BishopMoves(x, y int) (moves []*Move) {
	// Get the figure
	figure := self.Get(x, y)
	if figure == nil || figure.model != Bishop {
		return
	}

	// Left and right
	for _, diagonal := range []int{-1, 1} {

		// Up and down
		for _, direction := range []int{-1, 1} {

			// The movement offsets
			offx, offy := direction*diagonal, direction

			// Walk into direction
			for nx, ny := x+offx, y+offy; nx >= 0 && nx < 8 &&
				ny >= 0 && ny < 8; nx, ny = nx+offx, ny+offy {

				// Check board bounds
				if _, _, ok := self.ValidateWalls(nx, ny); ok {
					// Get figure
					enemy := self.Get(nx, ny)

					// If there is no figure or the figure is an enemy
					if enemy == nil || enemy.white != figure.white {
						moves = append(moves, &Move{figure, enemy, x, y, nx, ny, false})
					}

					if enemy != nil {
						break
					}
				} else {
					break
				}
			}
		}
	}
	return
}

func (self *Board) QueenMoves(x, y int) (moves []*Move) {
	// Get the figure
	figure := self.Get(x, y)
	if figure == nil || figure.model != Queen {
		return
	}

	// Going vertically
	for _, direction := range []int{-1, 1} {

		// Walk into direction
		for ny := y + direction; ny >= 0 && ny < 8; ny += direction {

			// Check board bounds
			if _, _, ok := self.ValidateWalls(x, ny); ok {
				// Get figure
				enemy := self.Get(x, ny)

				// If there is no figure or the figure is an enemy
				if enemy == nil || enemy.white != figure.white {
					moves = append(moves, &Move{figure, enemy, x, y, x, ny, false})
				}

				if enemy != nil {
					break
				}
			} else {
				break
			}
		}
	}

	// Going horizontally
	for _, direction := range []int{-1, 1} {
		// Walk into direction
		for nx := x + direction; nx >= 0 && nx < 8; nx += direction {
			// Check board bounds
			if _, _, ok := self.ValidateWalls(nx, y); ok {
				// Get figure
				enemy := self.Get(nx, y)

				// If there is no figure or the figure is an enemy
				if enemy == nil || enemy.white != figure.white {
					moves = append(moves, &Move{figure, enemy, x, y, nx, y, false})
				}

				if enemy != nil {
					break
				}
			} else {
				break
			}
		}
	}

	// Left and right
	for _, diagonal := range []int{-1, 1} {

		// Up and down
		for _, direction := range []int{-1, 1} {

			// The movement offsets
			offx, offy := direction*diagonal, direction

			// Walk into direction
			for nx, ny := x+offx, y+offy; nx >= 0 && nx < 8 &&
				ny >= 0 && ny < 8; nx, ny = nx+offx, ny+offy {

				// Check board bounds
				if _, _, ok := self.ValidateWalls(nx, ny); ok {
					// Get figure
					enemy := self.Get(nx, ny)

					// If there is no figure or the figure is an enemy
					if enemy == nil || enemy.white != figure.white {
						moves = append(moves, &Move{figure, enemy, x, y, nx, ny, false})
					}

					if enemy != nil {
						break
					}
				} else {
					break
				}
			}
		}
	}
	return
}

func main() {
	var board Board
	board.Set(3, 3, &Figure{false, Rook})
	board.Set(1, 0, &Figure{true, Knight})
	board.Set(1, 3, &Figure{true, Queen})
	board.Set(1, 1, &Figure{true, Bishop})
	board.Set(2, 1, &Figure{true, Pawn})
	board.Set(2, 3, &Figure{false, Pawn})
	board.Set(1, 2, &Figure{false, Pawn})

	for _, row := range board {
		for _, piece := range row {
			if piece != nil {
				fmt.Print(piece.model)
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}

	for _, i := range board.QueenMoves(1, 3) {
		fmt.Println(i)
	}
}
