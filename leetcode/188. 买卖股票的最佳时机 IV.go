package main

// 188. 买卖股票的最佳时机 IV

// 给定一个整数数组prices ，它的第 i 个元素prices[i] 是一支给定的股票在第 i 天的价格。
//
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

func main() {

}

// 123的进阶版，当k==2时为123的题意
//
// 一天一共就有2*k+1个状态：(可以发现，奇数都是买入、偶数都是卖出)
// 0. 没有操作
// 1. 第一次买入
// 2. 第一次卖出
// 3. 第二次买入
// 4. 第二次卖出
// 5. 第三次买入
// 6. 第三次卖出
// .......
func maxProfit3(k int, prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}

	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2*k+1)
	}

	// 初始化dp数组
	dp[0][0] = 0
	for i := 1; i < 2*k; i += 2 {
		// 当i为奇数时初始化第i次买入的剩余现金数、当i为偶数时剩余现金数初始化为0
		dp[0][i] = -prices[0]
	}

	// 遍历数组
	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		for j := 1; j <= 2*k; j++ {
			// 1. 保持前一天的未操作状态
			dp[i][0] = dp[i-1][0]
			if j%2 != 0 {
				// 2. 奇数为买入状态
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i]) // 保持前一天的买入状态或在前一天卖出的情况下买入
			} else {
				// 3. 偶数为卖出状态
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+prices[i]) // 保持前一天的卖出状态或在前一天没有卖出的情况下卖出
			}
		}
	}

	return dp[len(prices)-1][2*k]
}
