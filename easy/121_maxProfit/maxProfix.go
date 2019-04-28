package maxProfit

import "math"

func maxProfit(prices []int) int {
	max := 0
	min := math.MaxInt64

	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > max {
			max = prices[i] - min
		}
	}

	return max

}
