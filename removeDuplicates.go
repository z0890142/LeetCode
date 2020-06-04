package main

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		nums[index] = nums[i]
		index++
	}
	return index
}
