package main

// 剑指 Offer 63. 股票的最大利润

// 假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？

func main() {

}

func maxProfit(prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	minCost, maxProfit := prices[0], 0 // 记录最小值和最大利润
	for i := 1; i < len(prices); i++ {
		if prices[i] < minCost { // 若当前值比最小值小，则更新最小值
			minCost = prices[i]
		} else {
			maxProfit = max(maxProfit, prices[i]-minCost) // 尝试更新最大利润
		}
	}

	return maxProfit
}
