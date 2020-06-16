package main

func generateMatrix(n int) [][]int {
	var result [][]int
	sum := 1
	start, end := 0, n-1
	for time := 0; time < n; time++ {
		tmp := make([]int, n)
		result = append(result, tmp)
	}
	for start < end {
		for count := start; count < end; count++ {
			result[start][count] = sum
			sum += 1
		}
		for count := start; count < end; count++ {
			result[count][end] = sum
			sum += 1
		}
		for count := end; count > start; count-- {
			result[end][count] = sum
			sum += 1
		}
		for count := end; count > start; count-- {
			result[count][start] = sum
			sum += 1
		}
		start += 1
		end -= 1
	}
	if n%2 > 0 {
		result[start][end] = sum
	}
	return result
}
