package main

import "math"

func reverse(x int) int {
	result := 0

	for x != 0 {

		tmpVal := x % 10
		result = result*10 + tmpVal
		x = (x - tmpVal) / 10
		if result > math.MaxInt32 || result < math.MinInt32 {
			return 0
		}
	}
	return result
}
