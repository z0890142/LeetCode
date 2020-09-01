package main

func uniquePaths(m int, n int) int {
	var result [][]int
	for mIndex := 0; mIndex < m; mIndex++ {
		tmp := make([]int, n)
		result = append(result, tmp)
	}
	for i := 0; i < m; i++ {
		result[i][0] = 1
	}
	for i := 0; i < n; i++ {
		result[0][i] = 1
	}
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			result[row][col] = result[row-1][col] + result[row][col-1]
		}
	}

	return result[m-1][n-1]
}
