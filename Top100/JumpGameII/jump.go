package JumpGameII

func jump(nums []int) int {
	l := len(nums)
	if l == 0 || l == 1 {
		return 0
	}

	result := 0
	index := 0
	maxStep := 0

	for maxStep < l-1 {
		result++
		pre := maxStep
		for ; index <= pre; index++ {
			if nums[index]+index > maxStep {
				maxStep = nums[index] + index
			}
		}
		if pre == maxStep {
			return -1
		}
	}

	return result
}
