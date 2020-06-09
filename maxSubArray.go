package main

import "math"

func maxSubArray(nums []int) int {
	var result int
	result = math.MinInt64
	sum := 0
	for _, n := range nums {
		if sum < 0 {
			sum = n
		} else {
			sum += n
		}
		if sum > result {
			result = sum
		}
	}
	return result
}
