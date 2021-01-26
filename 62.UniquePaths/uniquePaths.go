package UniquePaths

func uniquePaths(m int, n int) int {
	var dp [][]int
	row := m
	col := n

	for rowIndex := 0; rowIndex < row; rowIndex++ {
		tmp := make([]int, col)
		dp = append(dp, tmp)
	}

	for rowIndex := 0; rowIndex < row; rowIndex++ {
		dp[rowIndex][0] = 1
	}
	for colIndex := 0; colIndex < col; colIndex++ {
		dp[0][colIndex] = 1
	}
	for rowIndex := 1; rowIndex < row; rowIndex++ {
		for colIndex := 1; colIndex < col; colIndex++ {

			dp[rowIndex][colIndex] = dp[rowIndex-1][colIndex] + dp[rowIndex][colIndex-1]
		}
	}
	return dp[row-1][col-1]

}
