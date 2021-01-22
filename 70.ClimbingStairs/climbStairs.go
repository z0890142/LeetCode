package ClimbingStairs

func climbStairs(n int) int {
	dp := make([]int, n+1)
	if n == 1 {
		return 1
	}
	dp[1] = 1
	dp[2] = 2
	for index := 3; index <= n; index++ {
		dp[index] = dp[index-1] + dp[index-2]
	}

	return dp[n]
}
