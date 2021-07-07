package main

// 122. 买卖股票的最佳时机 II

// 给定一个数组 prices ，其中prices[i] 是一支给定股票第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

func main() {

}

// 等价于每天都买卖 -- 前一天比后一天小,买卖; 前一天比后一天大,不买卖
func maxProfit1(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}

	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}
