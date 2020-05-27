package main

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	var result int
	diff := 65535
	sort.Ints(nums)
	for index := 0; index < len(nums); index++ {
		left, right := index+1, len(nums)-1
		var tmpDiff int
		for left < right {
			sum := nums[index] + nums[left] + nums[right]
			if sum < target {
				tmpDiff = target - sum
				left++
			} else {
				tmpDiff = sum - target
				right--
			}
			if tmpDiff == 0 {
				return sum
			}

			if tmpDiff < diff {
				diff = tmpDiff
				result = sum
			}
		}
	}
	return result
}
