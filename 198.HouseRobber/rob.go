package HouseRobber

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))

	if len(nums) == 1 {
		return nums[0]
	}

	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for index := 2; index < len(nums); index++ {
		dp[index] = max(dp[index-2]+nums[index], dp[index-1])
	}
	return dp[len(nums)-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
