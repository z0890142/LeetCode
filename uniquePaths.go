package main

func uniquePaths(m int, n int) int {
	var result [][]int
	for mIndex := 0; mIndex < m; mIndex++ {
		tmp := make([]int, n)
		result = append(result, tmp)
	}
	for mIndex := 0; mIndex < m; mIndex++ {
		for nIndex := 0; nIndex < n; nIndex++ {
			if mIndex == 0 || nIndex == 0 {
				result[mIndex][nIndex] = 1
			} else {
				result[mIndex][nIndex] = result[mIndex-1][nIndex] + result[mIndex][nIndex-1]
			}
		}
	}
	return result[m-1][n-1]
}
