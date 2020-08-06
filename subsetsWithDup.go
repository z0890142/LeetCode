package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var result [][]int
	result = append(result, []int{})
	sort.Ints(nums)
	subsetsWithDupHelper(nums, []int{}, &result)
	return result
}

func subsetsWithDupHelper(nums []int, preSlices []int, result *[][]int) {

	for index := 0; index < len(nums); index++ {

		preSlices = append(preSlices, nums[index])
		tmp := make([]int, len(preSlices))
		copy(tmp, preSlices)
		*result = append(*result, tmp)

		subsetsWithDupHelper(nums[index+1:], preSlices, result)

		preSlices = preSlices[0 : len(preSlices)-1]

		for index+1 < len(nums) && nums[index] == nums[index+1] {
			index++
		}
	}
}
