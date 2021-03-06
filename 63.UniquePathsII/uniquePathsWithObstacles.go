package UniquePathsII

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var dp [][]int
	row := len(obstacleGrid)
	col := len(obstacleGrid[0])
	for rowIndex := 0; rowIndex < row; rowIndex++ {
		tmp := make([]int, col)
		dp = append(dp, tmp)
	}

	for rowIndex := 0; rowIndex < row; rowIndex++ {
		if obstacleGrid[rowIndex][0] == 1 {
			break
		}
		dp[rowIndex][0] = 1
	}
	for colIndex := 0; colIndex < col; colIndex++ {
		if obstacleGrid[0][colIndex] == 1 {
			break
		}
		dp[0][colIndex] = 1
	}

	for rowIndex := 1; rowIndex < row; rowIndex++ {
		for colIndex := 1; colIndex < col; colIndex++ {
			if obstacleGrid[rowIndex][colIndex] == 1 {
				continue
			}
			dp[rowIndex][colIndex] = dp[rowIndex-1][colIndex] + dp[rowIndex][colIndex-1]
		}
	}
	return dp[row-1][col-1]

}
