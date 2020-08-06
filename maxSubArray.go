package main

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	max := nums[0]
	dp[0] = nums[0]
	for index := 1; index < len(nums); index++ {
		if dp[index-1]+nums[index] < nums[index] {
			dp[index] = nums[index]
		} else {
			dp[index] = dp[index-1] + nums[index]
		}
		if dp[index] > max {
			max = dp[index]
		}
	}
	return max
}
