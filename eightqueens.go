package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	board := [8]int{}
	solveQueens(board, 0)
}

func solveQueens(board [8]int, col int) {
	if col == 8 {

		for _, row := range board {
			z01.PrintRune(rune(row + '0'))
		}
		z01.PrintRune('\n')
		return
	}

	for row := 1; row <= 8; row++ {
		if isSafe(board, col, row) {
			board[col] = row
			solveQueens(board, col+1)
		}
	}
}

func isSafe(board [8]int, col int, row int) bool {
	for c := 0; c < col; c++ {
		r := board[c]
		if r == row ||
			r-c == row-col ||
			r+c == row+col {
			return false
		}
	}
	return true
}
