package main

func maxProfit1(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}

	min := prices[0]
	maxProfit := 0
	for _, price := range prices {
		if price < min {
			min = price
		}
		profit := price - min
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}
