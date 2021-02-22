package ValidSudoku

func isValidSudoku(board [][]byte) bool {
	m := len(board)
	n := len(board[0])

	for i := 0; i < m; i++ {
		rowMap := map[byte]int{}
		colMap := map[byte]int{}
		cubeMap := map[byte]int{}

		for j := 0; j < n; j++ {
			if board[i][j] != '.' {
				if _, ok := rowMap[board[i][j]]; ok {
					return false
				}
				rowMap[board[i][j]]++
			}

			if board[j][i] != '.' {
				if _, ok := colMap[board[j][i]]; ok {
					return false
				}
				colMap[board[j][i]]++
			}

			cubeRowIndex := 3 * (i / 3)
			cubeColIndex := 3 * (i % 3)
			if board[cubeRowIndex+j/3][cubeColIndex+j%3] != '.' {
				if _, ok := cubeMap[board[cubeRowIndex+j/3][cubeColIndex+j%3]]; ok {
					return false
				}
				cubeMap[board[cubeRowIndex+j/3][cubeColIndex+j%3]]++
			}

		}
	}
	return true

}
