package main

// 309. 最佳买卖股票时机含冷冻期

// 给定一个整数数组，其中第i个元素代表了第i天的股票价格。
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
//
// 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。

// func main() {

// }

// 类似 offer 63,Hot100 121
func maxProfit2(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	cost := -1     // 买入的成本 -- 若当前没有交易则为-1
	maxProfit := 0 // 最大利润

	for i := 0; i < len(prices); i++ {
		if i == len(prices)-1 {

		}

		if prices[i] > prices[i+1] {

		}
		if prices[i] > cost {
			maxProfit = max(maxProfit, prices[i]-cost)
		} else {
			cost = prices[i]
		}
	}

	return maxProfit
}
