package main

func minPathSum(grid [][]int) int {
	var result [][]int
	row := len(grid)
	col := len(grid[0])
	for rowIndex := 0; rowIndex < row; rowIndex++ {
		tmp := make([]int, col)
		result = append(result, tmp)
	}

	result[0][0] = grid[0][0]
	for rowIndex := 1; rowIndex < row; rowIndex++ {
		result[rowIndex][0] = result[rowIndex-1][0] + grid[rowIndex][0]
	}
	for colIndex := 1; colIndex < col; colIndex++ {

		result[0][colIndex] = result[0][colIndex-1] + grid[0][colIndex]

	}

	for rowIndex := 1; rowIndex < row; rowIndex++ {
		for colIndex := 1; colIndex < col; colIndex++ {
			tmp1 := result[rowIndex][colIndex-1] + grid[rowIndex][colIndex]
			tmp2 := result[rowIndex-1][colIndex] + grid[rowIndex][colIndex]
			if tmp1 <= tmp2 {
				result[rowIndex][colIndex] = tmp1
				continue
			}
			result[rowIndex][colIndex] = tmp2
		}
	}
	return result[row-1][col-1]
}
