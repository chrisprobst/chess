package main

import (
	"strconv"
	"unicode"
)

func ParseCoords(column, row string) (int, int, error) {
	if x, ok := LetterToColumn(rune(column[0])); ok {
		if y, err := strconv.ParseInt(row, 10, 32); err == nil {
			return x, int(y), nil
		} else {
			return -1, -1, err
		}
	}
	return -1, -1, nil
}

func LetterToColumn(letter rune) (int, bool) {
	if x := int(unicode.ToLower(letter) - 'a'); x >= 0 && x <= 7 {
		return 7 - x, true
	}
	return -1, false
}
