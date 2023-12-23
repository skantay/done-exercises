package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	n := 8
	solveQueens(n)
}

func solveQueens(n int) {
	board := []int{-1, -1, -1, -1, -1, -1, -1, -1} // creating a board 8x8

	solveQueensUtil(board, 0, n)
}

func isSafe(board []int, row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i] == col || board[i]-i == col-row || board[i]+i == col+row {
			return false
		}
	}
	return true
}

func printSolution(board []int) {
	for _, col := range board {
		z01.PrintRune(rune(col + 1 + 48))
	}
	z01.PrintRune(10)
}

func solveQueensUtil(board []int, row, n int) {
	if row == n {
		printSolution(board)
		return
	}

	for col := 0; col < n; col++ {
		if isSafe(board, row, col) {
			board[row] = col
			solveQueensUtil(board, row+1, n)
			board[row] = -1
		}
	}
}
