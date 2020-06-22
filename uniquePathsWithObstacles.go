package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row := len(obstacleGrid)
	col := len(obstacleGrid[0])
	var result [][]int

	for rowIndex := 0; rowIndex < row; rowIndex++ {
		tmp := make([]int, col)
		result = append(result, tmp)
	}
	if obstacleGrid[0][0] == 0 {
		result[0][0] = 1
	}

	for rowIndex := 1; rowIndex < row; rowIndex++ {
		if obstacleGrid[rowIndex][0] == 0 {
			result[rowIndex][0] = result[rowIndex-1][0]
		}
	}

	for colIndex := 1; colIndex < col; colIndex++ {
		if obstacleGrid[0][colIndex] == 0 {
			result[0][colIndex] = result[0][colIndex-1]
		}
	}

	for rowIndex := 1; rowIndex < row; rowIndex++ {
		for colIndex := 1; colIndex < col; colIndex++ {
			if obstacleGrid[rowIndex][colIndex] == 0 {
				result[rowIndex][colIndex] = result[rowIndex-1][colIndex] + result[rowIndex][colIndex-1]
			}
		}
	}
	return result[row-1][col-1]
}
