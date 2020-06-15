package main

func canJump(nums []int) bool {
	maxStep := 0
	for i := 0; i < len(nums); i++ {
		if i > maxStep {
			return false
		}
		if i+nums[i] > maxStep {
			maxStep = i + nums[i]
		}
	}
	return true
}
