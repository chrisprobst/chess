package main

import (
	"errors"
	"strconv"
	"unicode"
)

func ParseCoords(column, row string) (int, int, error) {
	// Check length of column
	if len(column) != 1 {
		return -1, -1, errors.New("Column string must have length of 1")
	}

	// Convert the letter to an index
	x, err := letterToColumn(rune(column[0]))
	if err != nil {
		return -1, -1, err
	}

	// Try to parse the row
	y, err := strconv.ParseInt(row, 10, 32)
	if err != nil {
		return -1, -1, err
	}

	if y < 1 || y > 8 {
		return -1, -1, errors.New("Row index must be between 1 and 8")
	}

	// Return both values
	return x, int(y) - 1, nil
}

func letterToColumn(letter rune) (int, error) {
	if x := int(unicode.ToLower(letter) - 'a'); x >= 0 && x <= 7 {
		return 7 - x, nil
	}
	return -1, errors.New("Letter not within 'a' - 'H'")
}
