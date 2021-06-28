package main

// 121. 买卖股票的最佳时机

// 给定一个数组 prices ，它的第i 个元素prices[i] 表示一支给定股票第 i 天的价格。
// 你只能选择某一天买入这只股票，并选择在未来的某一个不同的日子卖出该股票。
// 设计一个算法来计算你所能获取的最大利润。
//
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

func main() {

}

func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	minCost := prices[0] // 最小成本
	maxProfit := 0       // 最大利润
	for _, price := range prices {
		if price < minCost {
			minCost = price
		} else {
			maxProfit = max(maxProfit, price-minCost)
		}
	}

	return maxProfit
}
