package main

func isValidSudoku(board [][]string) bool {

	n := len(board)
	m := len(board[0])
	for rowIndex := 0; rowIndex < m; rowIndex++ {
		rows := map[string]int{}
		cols := map[string]int{}
		cube := map[string]int{}
		for colIndex := 0; colIndex < n; colIndex++ {
			if string(board[rowIndex][colIndex]) != "." {
				if _, ok := rows[board[rowIndex][colIndex]]; ok {
					return false
				}
				rows[board[rowIndex][colIndex]]++
			}

			if string(board[colIndex][rowIndex]) != "." {
				if _, ok := cols[board[colIndex][rowIndex]]; ok {
					return false
				}
				rows[board[colIndex][rowIndex]]++
			}

			cubeRowIndex := 3 * (rowIndex / 3)
			cubeColIndex := 3 * (rowIndex % 3)

			if board[cubeRowIndex+colIndex/3][cubeColIndex+colIndex%3] != "." {
				if _, ok := cube[board[cubeRowIndex+colIndex/3][cubeColIndex+colIndex%3]]; ok {
					return false
				}
				cube[board[cubeRowIndex+colIndex/3][cubeColIndex+colIndex%3]]++
			}

		}

	}
	return true

}
