package maxProfit

func maxProfit(prices []int) int {
	buy := 0
	sell := 0
	profit := 0

	n := len(prices)
	for {
		for buy < n-1 && prices[buy+1] <= prices[buy] {
			buy++
		}

		sell = buy

		for sell < n-1 && prices[sell+1] > prices[sell] {
			sell++
		}

		if sell == buy {
			return profit
		}

		profit += prices[sell] - prices[buy]
		buy = sell
	}
}
