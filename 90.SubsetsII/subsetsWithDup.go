package SubsetsII

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)

	result = append(result, []int{})
	subsetsWithDupHelper(nums, []int{}, &result)
	return result
}

func subsetsWithDupHelper(nums []int, tmp []int, result *[][]int) {
	for index, val := range nums {

	}
}
