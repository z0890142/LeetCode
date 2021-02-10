package SpiralMatrix

func spiralOrder(matrix [][]int) []int {
	var result []int
	if len(matrix) == 0 {
		return result
	}
	maxRow := len(matrix) - 1
	maxCol := len(matrix[0]) - 1
	var layer int

	if maxCol > maxRow {
		layer = maxRow / 2
	} else {
		layer = maxCol / 2
	}

	for count := 0; count <= layer; count++ {
		if count == maxRow-count {
			for i := count; i <= maxCol-count; i++ {
				result = append(result, matrix[count][i])
			}
		} else if count == maxCol-count {
			for i := count; i <= maxRow-count; i++ {
				result = append(result, matrix[i][count])
			}
		} else {
			for i := count; i < maxCol-count; i++ {
				result = append(result, matrix[count][i])
			}
			for i := count; i < maxRow-count; i++ {
				result = append(result, matrix[i][maxCol-count])
			}
			for i := count; i < maxCol-count; i++ {
				result = append(result, matrix[maxRow-count][maxCol-i])
			}
			for i := count; i < maxRow-count; i++ {
				result = append(result, matrix[maxRow-i][count])
			}
		}
	}
	return result
}
