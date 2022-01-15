package main

// 121. 买卖股票的最佳时机

// func main() {

// }

func maxProfit(prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	minCost, maxProfit := prices[0], 0 // 最小成本、最大利润
	for i := 1; i < len(prices); i++ {
		if prices[i] > minCost {
			maxProfit = max(maxProfit, prices[i]-minCost)
		} else {
			minCost = prices[i]
		}
	}

	return maxProfit
}
