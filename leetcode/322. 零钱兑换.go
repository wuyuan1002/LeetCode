package main

import "math"

// 322. 零钱兑换

// func main() {

// }

// 类似Hot100 518, 416，39(与回溯法的区别)
// 动态规划 -- dp[i]表示i元的最少硬币数 -- dp[i] = min(dp[i], dp[i-num]+1)
func coinChange(coins []int, amount int) int {
	if coins == nil || len(coins) == 0 {
		return -1
	}

	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}

	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for _, coin := range coins {
		for i := 1; i <= amount; i++ {
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
