package Searcha2DMatrix

func searchMatrix(matrix [][]int, target int) bool {

	for _, row := range matrix {
		min := row[0]
		max := row[len(row)-1]
		if target < min {
			return false
		}
		if target > max {
			continue
		}
		if target == min || target == max {
			return true
		}
		l, r := 0, len(row)-1
		for l <= r {
			mid := l + (r - l)
			if row[mid] > target {
				r--
			} else if row[mid] < target {
				l++
			} else {
				return true
			}
		}
		return false
	}
	return false
}
