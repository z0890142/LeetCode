package main

func removeDuplicates2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	target := nums[0]
	index := 1
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == target {
			if count < 2 {
				nums[index] = target
				index++
			}
			count++
			continue
		}
		nums[index] = nums[i]
		target = nums[i]
		count = 1
		index++
	}

	return index
}
