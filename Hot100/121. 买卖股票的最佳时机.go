package main

// 121. 买卖股票的最佳时机

// 给定一个数组 prices ，它的第i 个元素prices[i] 表示一支给定股票第 i 天的价格。
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。
// 设计一个算法来计算你所能获取的最大利润。
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

func main() {

}

// 同offer 63
func maxProfit(prices []int) int {
	if prices == nil {
		return 0
	}

	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	profit := 0          // 最大利润
	minCost := prices[0] // 最小成本
	r := 0
	for ; r < len(prices); r++ {
		// 更新最小成本
		minCost = min(minCost, prices[r])
		// 更新最大利润
		profit = max(profit, prices[r]-minCost)
	}

	return profit
}
