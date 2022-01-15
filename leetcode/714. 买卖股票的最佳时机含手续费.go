package main

// 714. 买卖股票的最佳时机含手续费

// 给定一个整数数组prices，其中第i个元素代表了第i天的股票价格 ；整数fee 代表了交易股票的手续费用。
//
// 你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，
// 在卖出它之前你就不能再继续购买股票了。
// 返回获得利润的最大值。
//
// 注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。

// func main() {
// 	fmt.Println(maxProfit5([]int{1, 3, 2, 8, 4, 9}, 2))
// }

// 同122
// 只是在卖出股票时需要同时扣除手续费
func maxProfit5(prices []int, fee int) int {
	if prices == nil || len(prices) == 0 || fee < 0 {
		return 0
	}

	dp := make([][2]int, len(prices))

	// 初始化dp
	dp[0][0] = -prices[0]
	dp[0][1] = 0

	// 遍历数组
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}

	return dp[len(prices)-1][1]
}
