package BestTimetoBuyandSellStock

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	maxProfit := 0
	min := prices[0]
	for _, val := range prices {
		if val < min {
			min = val
			continue
		}
		profit := val - min
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}
