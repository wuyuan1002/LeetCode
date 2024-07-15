package main

// 121. 买卖股票的最佳时机

// 给定一个数组 prices ，它的第i 个元素prices[i] 表示一支给定股票第 i 天的价格。
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。
// 设计一个算法来计算你所能获取的最大利润。
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

// maxProfit .
// 因为股票只能买卖一次，那么贪心的想法很自然就是取左侧最小值，取右侧最大值，两者的差值就是最大利润
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	// 最小成本、最大利润
	minCost, result := prices[0], 0

	// 遍历每天的股票，更新最小成本并计算最大利润
	for _, price := range prices {
		if price < minCost {
			minCost = price
		} else {
			result = max(result, price-minCost)
		}
	}

	// 返回最大利润
	return result
}
