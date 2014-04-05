package api

func (self *board) pawnMoves(x, y int) (moves []*move) {

	// Get the figure
	fig := self.get(x, y)
	if fig == nil {
		return
	}

	// Color to direction
	direction := 1
	if !fig.white {
		direction = -1
	}

	// The pawn is at the end of the trail
	if _, _, ok := self.validateWalls(x, y+direction); !ok {
		return
	}

	// Check if this pawn would transform into something
	var transformInto []*figure
	if newLoc := y + direction; newLoc == 0 || newLoc == 7 {
		transformInto = append(transformInto,
			newFigure(fig.white, knight),
			newFigure(fig.white, bishop),
			newFigure(fig.white, rook),
			newFigure(fig.white, queen))
	}

	// Check usual forward movement
	if self.isFree(x, y+direction) {

		// Append the forward move
		moves = append(moves, &move{fig, nil, x, y, x, y + direction, transformInto, nil})

		// Check end of trail
		if _, _, ok := self.validateWalls(x, y+2*direction); ok {

			// Check double forward movement
			if (y == 1 || y == 6) && self.isFree(x, y+2*direction) {

				// Append double move
				moves = append(moves, &move{fig, nil, x, y, x, y + 2*direction, nil, nil})
			}
		}
	}

	// Check if we can hit an enemy
	for _, offset := range []int{-1, 1} {
		if enemy := self.get(x+offset, y+direction); enemy != nil && enemy.white != fig.white {
			moves = append(moves, &move{fig, enemy, x, y, x + offset, y + direction, transformInto, nil})
		}
	}

	return
}

func (self *board) knightMoves(x, y int) (moves []*move) {
	// Get the figure
	fig := self.get(x, y)
	if fig == nil {
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
			if _, _, ok := self.validateWalls(nx, ny); ok {

				// Get the figure at this location
				enemy := self.get(nx, ny)

				// If there is no figure or the figure is an enemy
				if enemy == nil || enemy.white != fig.white {

					// Append this move
					moves = append(moves, &move{fig, enemy, x, y, nx, ny, nil, nil})
				}
			}
		}
	}

	return
}

func (self *board) straightMoves(x, y, maxHops int) (moves []*move) {
	// Get the figure
	fig := self.get(x, y)
	if fig == nil {
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
				if _, _, ok := self.validateWalls(nx, ny); ok {
					// Get figure
					enemy := self.get(nx, ny)

					// If there is no figure or the figure is an enemy
					if enemy == nil || enemy.white != fig.white {
						moves = append(moves, &move{fig, enemy, x, y, nx, ny, nil, nil})
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

func (self *board) diagonalMoves(x, y, maxHops int) (moves []*move) {
	// Get the figure
	fig := self.get(x, y)
	if fig == nil {
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
				if _, _, ok := self.validateWalls(nx, ny); ok {
					// Get figure
					enemy := self.get(nx, ny)

					// If there is no figure or the figure is an enemy
					if enemy == nil || enemy.white != fig.white {
						moves = append(moves, &move{fig, enemy, x, y, nx, ny, nil, nil})
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

func (self *board) rookMoves(x, y int) []*move {
	return self.straightMoves(x, y, 8)
}

func (self *board) bishopMoves(x, y int) []*move {
	return self.diagonalMoves(x, y, 8)
}

func (self *board) queenMoves(x, y int) []*move {
	return append(self.rookMoves(x, y), self.bishopMoves(x, y)...)
}

func (self *board) kingMoves(x, y int, ignoreCheckAndCastling bool) (moves []*move) {
	// Get the figure
	figure := self.get(x, y)
	if figure == nil {
		return
	}

	// Usual movement of the king
	tmpMoves := self.diagonalMoves(x, y, 1)
	tmpMoves = append(tmpMoves, self.straightMoves(x, y, 1)...)

	// Iterate over all tmp moves
	for _, move := range tmpMoves {
		if ignoreCheckAndCastling || self.isValidKingMove(move) {
			moves = append(moves, move)
		}
	}

	// The king-rook moves
	if !ignoreCheckAndCastling && !self.isKingChecked(x, y) && figure.flags&moved == 0 {

		// For both rooks
		for dx := range []int{0, 7} {

			// Lookup rook
			rook := self.get(dx, y)

			// If rook exists and has the same color and has not moved yet!
			if rook != nil && rook.white == figure.white && rook.flags&moved == 0 {

				// Compute the relevant scalars
				offset, direction, distance := 1, 1, x-2
				if dx == 7 {
					offset, direction, distance = 6, -1, 7-x-2
				}

				// Check if the rook can move
				freeSpace := true
				for i := 0; i < distance; i++ {
					if !self.isFree(offset+i*direction, y) {
						freeSpace = false
						break
					}
				}

				if freeSpace {
					// Create the king-rook move
					kingRookMove := &move{figure, nil, x, y, x - 2*direction, y, nil,
						&move{rook, nil, 0, y, x - direction, y, nil, nil},
					}

					// Add the king-rook move if valid
					if self.isValidKingMove(kingRookMove) {
						moves = append(moves, kingRookMove)
					}
				}
			}
		}
	}

	return
}

func (self *board) movesExt(x, y int, ignoreCheckAndCastling bool) (moves []*move) {
	fig := self.get(x, y)
	if fig == nil {
		return nil
	}

	switch fig.model {
	case pawn:
		moves = self.pawnMoves(x, y)
	case knight:
		moves = self.knightMoves(x, y)
	case rook:
		moves = self.rookMoves(x, y)
	case bishop:
		moves = self.bishopMoves(x, y)
	case queen:
		moves = self.queenMoves(x, y)
	case king:
		moves = self.kingMoves(x, y, ignoreCheckAndCastling)
	}
	return
}

func (self *board) moves(x, y int) []*move {
	return self.movesExt(x, y, false)
}
