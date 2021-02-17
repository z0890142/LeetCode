package PermutationsII

import (
	"sort"
)

func permuteUnique(nums []int) [][]int {
	resp := [][]int{}

	sort.Ints(nums)
	bt(&resp, []int{}, nums)

	return resp
}

func bt(resp *[][]int, sub []int, nums []int) {
	if len(nums) == 0 {
		*resp = append(*resp, append([]int{}, sub...))
		return
	}

	for i := range nums {
		if i == 0 || nums[i] != nums[i-1] {
			nums2 := append(append([]int{}, nums[0:i]...), nums[i+1:]...)
			bt(resp, append(sub, nums[i]), nums2)
		}
	}
}
