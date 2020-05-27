package main

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	lIndex := 0
	for index := 1; index < len(nums); index++ {
		if nums[index-1] != nums[index] {
			nums[lIndex+1] = nums[index]
			lIndex += 1
		}
	}
	return lIndex + 1
}
