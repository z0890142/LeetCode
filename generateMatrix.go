package main

func generateMatrix(n int) [][]int {
	var result [][]int
	if n == 0 {
		return result
	}
	for time := 0; time < n; time++ {
		tmp := make([]int, n)
		result = append(result, tmp)
	}
	num := 1
	for layout := 0; layout <= n/2; layout++ {
		if layout == n/2 && n%2 != 0 {
			result[layout][layout] = num
			break
		}
		for i := layout; i < n-layout-1; i++ {
			result[layout][i] = num
			num++
		}
		for i := layout; i < n-layout-1; i++ {
			result[i][n-layout-1] = num
			num++
		}
		for i := n - layout - 1; i > layout; i-- {
			result[n-layout-1][i] = num
			num++
		}
		for i := n - layout - 1; i > layout; i-- {
			result[i][layout] = num
			num++
		}
	}
	return result
}
