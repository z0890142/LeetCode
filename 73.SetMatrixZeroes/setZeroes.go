package SetMatrixZeroes

func setZeroes(matrix [][]int) [][]int {

	colZero := make([]bool, len(matrix[0]))
	rowZero := make([]bool, len(matrix))

	for rowIndex, row := range matrix {
		for colIndex, col := range row {
			if col == 0 {
				colZero[colIndex] = true
				rowZero[rowIndex] = true
			}
		}
	}

	for index, val := range colZero {
		if val {
			for row, _ := range matrix {
				matrix[row][index] = 0
			}
		}

	}

	for index, val := range rowZero {
		if val {
			for col, _ := range matrix[index] {
				matrix[index][col] = 0
			}
		}

	}

	return matrix

}
