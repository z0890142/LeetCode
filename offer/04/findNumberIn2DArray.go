package findNumberIn2DArray

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	var row = len(matrix[0]) - 1
	i := 0
	for row >= 0 && i < len(matrix) {
		if matrix[i][row] < target {
			i++
		} else if matrix[i][row] > target {
			row--
		} else {
			return true
		}
	}
	return false
}
