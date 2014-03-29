package main

func (self *Board) IsFoundKingChecked(white bool) bool {
	if x, y, ok := self.Find(white, King); ok {
		return self.IsKingChecked(x, y)
	}
	return false
}

func (self *Board) IsKingChecked(x, y int) bool {
	// Get the king
	king := self.Get(x, y)
	if king == nil || king.Model != King {
		return false
	}

	// Check if any enemy threaten the king
	for ny := 0; ny < 8; ny++ {
		for nx := 0; nx < 8; nx++ {
			if enemy := self.Get(nx, ny); enemy != nil && enemy.White != king.White {
				for _, move := range self.moves(nx, ny, true) {
					if move.FigureDest == king {
						return true
					}
				}
			}
		}
	}

	return false
}

func (self *Board) IsValidKingMove(move *Move) bool {
	return !self.ApplyMove(move).IsFoundKingChecked(move.FigureOrigin.White)
}
