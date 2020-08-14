package main

func canJump(nums []int) bool {
	maxStep := nums[0]
	for index := 1; index < len(nums); index++ {
		if index > maxStep {
			return false
		}
		if nums[index]+index > maxStep {
			maxStep = nums[index] + index
		}
	}
	return true
}
