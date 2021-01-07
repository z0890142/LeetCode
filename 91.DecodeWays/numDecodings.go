package DecodeWays

import "strings"

//如果說i-1不等於0代表我們是在i-1種解碼方式上在進行解碼
//如果i-2跟i-1可以組成10-26代表我們是在i-2種解碼方式上進行解碼
func numDecodings(s string) int {
	if len(s) == 0 || strings.HasPrefix(s, "0") {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}
		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6') {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}
