package main

func (self *Board) PawnMoves(x, y int) (moves []*Move) {

	// Get the figure
	figure := self.Get(x, y)
	if figure == nil {
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
	transformMove := NoTransform
	if newLoc := y + direction; newLoc == 0 || newLoc == 7 {
		transformMove = PawnTransform
	}

	// Check usual forward movement
	if self.IsFree(x, y+direction) {

		// Append the forward move
		moves = append(moves, &Move{figure, nil, x, y, x, y + direction, transformMove})

		// Check double forward movement
		if (y == 1 || y == 6) && self.IsFree(x, y+2*direction) {

			// Append double move
			moves = append(moves, &Move{figure, nil, x, y, x, y + 2*direction, NoTransform})
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
	if figure == nil {
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
					moves = append(moves, &Move{figure, enemy, x, y, nx, ny, NoTransform})
				}
			}
		}
	}

	return
}

func (self *Board) straightMoves(x, y, maxHops int) (moves []*Move) {
	// Get the figure
	figure := self.Get(x, y)
	if figure == nil {
		return
	}

	// Iterate over axes
	for _, axis := range []int{1, 0} {

		// Iterate over directions
		for _, direction := range []int{-1, 1} {

			// Walk into direction
			for offset := direction; offset >= -maxHops && offset <= maxHops; offset += direction {

				// Convert axis and direction to nx and ny
				nx, ny := x+axis*offset, y+(1-axis)*offset

				// Check board bounds
				if _, _, ok := self.ValidateWalls(nx, ny); ok {
					// Get figure
					enemy := self.Get(nx, ny)

					// If there is no figure or the figure is an enemy
					if enemy == nil || enemy.white != figure.white {
						moves = append(moves, &Move{figure, enemy, x, y, nx, ny, NoTransform})
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

func (self *Board) diagonalMoves(x, y, maxHops int) (moves []*Move) {
	// Get the figure
	figure := self.Get(x, y)
	if figure == nil {
		return
	}

	// Left and right
	for _, diagonal := range []int{-1, 1} {

		// Up and down
		for _, direction := range []int{-1, 1} {

			// The movement offsets
			dx, dy := direction*diagonal, direction

			// Walk into direction
			for offX, offY := dx, dy; offX >= -maxHops && offX <= maxHops &&
				offY >= -maxHops && offY <= maxHops; offX, offY = offX+dx, offY+dy {

				// Absolute position
				nx, ny := x+offX, y+offY

				// Check board bounds
				if _, _, ok := self.ValidateWalls(nx, ny); ok {
					// Get figure
					enemy := self.Get(nx, ny)

					// If there is no figure or the figure is an enemy
					if enemy == nil || enemy.white != figure.white {
						moves = append(moves, &Move{figure, enemy, x, y, nx, ny, NoTransform})
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

func (self *Board) RookMoves(x, y int) []*Move {
	return self.straightMoves(x, y, 8)
}

func (self *Board) BishopMoves(x, y int) []*Move {
	return self.diagonalMoves(x, y, 8)
}

func (self *Board) QueenMoves(x, y int) []*Move {
	return append(self.RookMoves(x, y), self.BishopMoves(x, y)...)
}

func (self *Board) KingMoves(x, y int) (moves []*Move) {
	// Get the figure
	figure := self.Get(x, y)
	if figure == nil {
		return
	}

	// Usual movement of the king
	tmpMoves := self.diagonalMoves(x, y, 1)
	tmpMoves = append(tmpMoves, self.straightMoves(x, y, 1)...)

	// Iterate over all tmp moves
	for _, move := range tmpMoves {

		// Create a new board for each move
		copy := self.ApplyMove(move)

		// The move does not provoke check
		if !copy.IsKingChecked(move.XDest, move.YDest) {
			moves = append(moves, move)
		}
	}

	return
}

func (self *Board) Moves(x, y int) (moves []*Move) {
	figure := self.Get(x, y)
	if figure == nil {
		return nil
	}

	switch figure.model {
	case Pawn:
		moves = self.PawnMoves(x, y)
	case Knight:
		moves = self.KnightMoves(x, y)
	case Rook:
		moves = self.RookMoves(x, y)
	case Bishop:
		moves = self.BishopMoves(x, y)
	case Queen:
		moves = self.QueenMoves(x, y)
	case King:
		moves = self.KingMoves(x, y)
	}
	return
}
