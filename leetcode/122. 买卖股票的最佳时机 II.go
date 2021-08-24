package main

// 122. 买卖股票的最佳时机 II

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
