package main

func setZeroes(matrix [][]int) [][]int {
	colZero := make([]int, len(matrix[0]))
	rowZero := make([]int, len(matrix))

	for rowIndex, row := range matrix {
		for colIndex, col := range row {
			if col == 0 {
				colZero[colIndex] = 1
				rowZero[rowIndex] = 1
			}
		}
	}

	for index, val := range colZero {
		if val == 0 {
			continue
		}
		for rowIndex, _ := range matrix {
			matrix[rowIndex][index] = 0
		}
	}
	for index, val := range rowZero {
		if val == 0 {
			continue
		}
		for colIndex, _ := range matrix[index] {
			matrix[index][colIndex] = 0
		}
	}
	return matrix
}
