package SurroundedRegions

import "fmt"

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	row := len(board) - 1
	col := len(board[0]) - 1
	for index, val := range board[0] {
		if val == 'O' {
			solveHelper(0, index, &board)
		}
	}
	for index, val := range board[row] {
		if val == 'O' {
			solveHelper(row, index, &board)
		}
	}
	for index := 0; index <= row; index++ {
		if board[index][0] == 'O' {
			solveHelper(index, 0, &board)
		}
	}
	for index := 0; index <= row; index++ {
		if board[index][col] == 'O' {
			solveHelper(index, col, &board)
		}
	}

	for rowIndex, col := range board {
		for colIndex, val := range col {
			if val == 'O' {
				board[rowIndex][colIndex] = 'X'
			} else if val == 'Y' {
				board[rowIndex][colIndex] = 'O'
			}
		}
	}
	fmt.Println(board)
}

func solveHelper(row int, col int, board *[][]byte) {
	(*board)[row][col] = 'Y'
	if row-1 > 0 && (*board)[row-1][col] == 'O' {
		solveHelper(row-1, col, board)
	}
	if row+1 < len(*board) && (*board)[row+1][col] == 'O' {
		solveHelper(row+1, col, board)
	}
	if col-1 > 0 && (*board)[row][col-1] == 'O' {
		solveHelper(row, col-1, board)
	}
	if col+1 < len((*board)[0]) && (*board)[row][col+1] == 'O' {
		solveHelper(row, col+1, board)
	}
}
