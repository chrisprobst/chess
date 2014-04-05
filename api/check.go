package main

func (self *board) isFoundKingChecked(white bool) bool {
	if x, y, ok := self.find(white, king); ok {
		return self.isKingChecked(x, y)
	}
	return false
}

func (self *board) isKingChecked(x, y int) bool {
	// Get the king
	kingFigure := self.get(x, y)
	if kingFigure == nil || kingFigure.model != king {
		return false
	}

	// Check if any enemy threaten the king
	for ny := 0; ny < 8; ny++ {
		for nx := 0; nx < 8; nx++ {
			if enemy := self.get(nx, ny); enemy != nil && enemy.white != kingFigure.white {
				for _, move := range self.movesExt(nx, ny, true) {
					if move.figureDest == kingFigure {
						return true
					}
				}
			}
		}
	}

	return false
}

func (self *board) isValidKingMove(move *move) bool {
	return !self.applyMove(move).isFoundKingChecked(move.figureOrigin.white)
}
