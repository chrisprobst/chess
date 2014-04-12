package rules

import (
	"fmt"
	"testing"
)

func TestApi(t *testing.T) {
	chess := NewChess()
	chess.Print()
	for _, move := range chess.Moves(1, 1) {
		fmt.Println(move)
	}
}
