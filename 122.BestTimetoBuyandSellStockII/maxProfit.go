package BestTimetoBuyandSellStockII

func maxProfit(prices []int) int {
	var total int
	if len(prices) == 0 {
		return 0
	}

	buy := prices[0]
	for index := 1; index < len(prices); index++ {
		if prices[index] > buy {
			total += prices[index] - buy
		}
		buy = prices[index]
	}
	return total
}
