package Triangle

import "sort"

func minimumTotal(triangle [][]int) int {
	maxRow := len(triangle) - 1
	for row := 1; row < len(triangle); row++ {
		for col := 0; col < len(triangle[row]); col++ {
			if col == 0 {
				triangle[row][col] = triangle[row][col] + triangle[row-1][col]
			} else if col == len(triangle[row])-1 {
				triangle[row][col] = triangle[row][col] + triangle[row-1][col-1]
			} else {
				triangle[row][col] = triangle[row][col] + min(triangle[row-1][col], +triangle[row-1][col-1])
			}
		}
	}
	sort.Ints(triangle[maxRow])
	return triangle[maxRow][0]
}

func min(a int, b int) int {
	if a > b {
		return b

	}
	return a
}
