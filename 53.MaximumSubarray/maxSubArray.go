package MaximumSubarray

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currSum := nums[0] // best sum including current number
	for _, n := range nums[1:] {
		currSum = max(n, currSum+n)
		maxSum = max(maxSum, currSum)
	}
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
