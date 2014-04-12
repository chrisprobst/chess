package rules

func (self *board) isFoundKingChecked(white bool) bool {
	if x, y, ok := self.find(white, king); ok {
		return self.isKingChecked(x, y)
	}
	return false
}

func (self *board) isKingCheckMate(x, y int) bool {
	kingFigure, checked := self.get(x, y), self.isKingChecked(x, y)
	if !checked {
		return false
	}

	for ny := 0; ny < 8; ny++ {
		for nx := 0; nx < 8; nx++ {
			if friend := self.get(nx, ny); friend != nil && friend.White == kingFigure.White {
				for _, move := range self.moves(x, y, false) {
					if !self.applyMove(move).isFoundKingChecked(kingFigure.White) {
						return false
					}
				}
			}
		}
	}
	return true
}

func (self *board) isKingChecked(x, y int) bool {
	// Get the king
	kingFigure := self.get(x, y)
	if kingFigure == nil || kingFigure.Model != king {
		return false
	}

	// Check if any enemy threaten the king
	for ny := 0; ny < 8; ny++ {
		for nx := 0; nx < 8; nx++ {
			if enemy := self.get(nx, ny); enemy != nil && enemy.White != kingFigure.White {
				for _, move := range self.moves(nx, ny, true) {
					if move.FigureDest == kingFigure {
						return true
					}
				}
			}
		}
	}

	return false
}

func (self *board) isValidKingMove(move *Move) bool {
	return !self.applyMove(move).isFoundKingChecked(move.FigureOrigin.White)
}
