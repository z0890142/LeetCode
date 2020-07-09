package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	colMax := len(matrix[0]) - 1
	for _, row := range matrix {
		if row[colMax] < target {
			continue
		}

		if row[0] > target {
			return false
		}
		if row[0] == target || row[colMax] == target {
			return true
		}
		l, r := 0, colMax
		for l <= r {
			mid := (l + r) / 2

			if row[mid] == target {
				return true
			} else if row[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		return false
	}
	return false
}
