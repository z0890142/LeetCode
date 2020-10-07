package main

func maxProfit(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}
	var result int
	min := prices[0]
	for _, val := range prices {
		if val < min {
			min = val
		}
		result += val - min
	}
	return result
}
