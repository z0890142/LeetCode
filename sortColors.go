package main

func sortColors(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	left, right := 0, length-1
	index := 0
	for index <= right {
		if nums[index] == 2 {
			nums[index], nums[right] = nums[right], nums[index]
			right -= 1
			continue
		} else if nums[index] == 0 {
			nums[index], nums[left] = nums[left], nums[index]
			left += 1
			index += 1
			continue
		} else {
			index += 1
			continue
		}
	}
}
