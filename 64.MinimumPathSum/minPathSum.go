package MinimumPathSum

func minPathSum(grid [][]int) int {
	var dp [][]int
	row := len(grid)
	col := len(grid[0])
	for rowIndex := 0; rowIndex < row; rowIndex++ {
		tmp := make([]int, col)
		dp = append(dp, tmp)
	}

	dp[0][0] = grid[0][0]

	for rowIndex := 1; rowIndex < row; rowIndex++ {
		dp[rowIndex][0] = dp[rowIndex-1][0] + grid[rowIndex][0]
	}
	for colIndex := 1; colIndex < col; colIndex++ {
		dp[0][colIndex] = dp[0][colIndex-1] + grid[0][colIndex]
	}

	for rowIndex := 1; rowIndex < row; rowIndex++ {
		for colIndex := 1; colIndex < col; colIndex++ {
			if dp[rowIndex][colIndex-1] < dp[rowIndex-1][colIndex] {
				dp[rowIndex][colIndex] = dp[rowIndex][colIndex-1] + grid[rowIndex][colIndex]
			} else {
				dp[rowIndex][colIndex] = dp[rowIndex-1][colIndex] + grid[rowIndex][colIndex]
			}
		}
	}

	return dp[row-1][col-1]

}
