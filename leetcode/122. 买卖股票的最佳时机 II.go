package main

// 122. 买卖股票的最佳时机 II

func main() {

}

// 1. 等价于每天都买卖 -- 前一天比后一天小,买卖; 前一天比后一天大,不买卖
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

// 2. 动态规划 -- 类似714
// 一天共有两种状态:
// 0: 持有股票
// 1: 未持有股票
//
// dp[i][j]表示第i天状态j时所持有的现金数量(越大说明越有钱)
func maxProfit11(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}

	dp := make([][2]int, len(prices))

	// 初始化dp
	dp[0][0] = -prices[0]
	dp[0][1] = 0

	// 遍历数组
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return dp[len(prices)-1][1]
}
