package SpiralMatrixII

func generateMatrix(n int) [][]int {
	var result [][]int

	if n == 0 {
		return result
	}

	for i := 0; i < n; i++ {
		result = append(result, make([]int, n))
	}

	total := 1
	for i := 0; i <= n/2; i++ {
		if i == n/2 && n%2 != 0 {
			result[i][i] = total
			break
		}
		for index := i; index < n-i-1; index++ {
			result[i][index] = total
			total++
		}
		for index := i; index < n-i-1; index++ {
			result[index][n-1-i] = total
			total++
		}
		for index := n - i - 1; index > i; index-- {
			result[n-1-i][index] = total
			total++
		}
		for index := n - i - 1; index > i; index-- {
			result[index][i] = total
			total++
		}
	}

	return result
}
