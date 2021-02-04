package HouseRobberII

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	return max(robHelper(nums, 0, len(nums)-1), robHelper(nums, 1, len(nums)))
}

func robHelper(nums []int, l int, r int) int {
	dp := make([]int, len(nums))
	dp[l] = nums[l]
	dp[l+1] = max(nums[l], nums[l+1])

	for index := l + 2; index < r; index++ {
		dp[index] = max(dp[index-1], dp[index-2]+nums[index])
	}
	return dp[r-1]

}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
