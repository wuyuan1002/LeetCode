package main

import (
	"fmt"
	"math"
)

// 322. 零钱兑换

// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
// 如果没有任何一种硬币组合能组成总金额，返回-1。
//
// 你可以认为每种硬币的数量是无限的。

func main() {
	fmt.Println(coinChange([]int{5, 2, 1}, 7))
}

// --------------------

// 回溯法和动态规划:
// 给定一个目标值，给定一个选择空间，问达到目标的解的个数 --(求解的个数、求是否存在问题、求排列组合)-- 回溯法
// 给定一个目标值，给定一个选择空间，问达到目标的最优解 --(求最优解、求是否存在问题)-- 背包问题

// --------------------

// 背包问题的转移方程:
// 1、最值问题: dp[i] = max/min(dp[i], dp[i-nums]+1)或dp[i] = max/min(dp[i], dp[i-num]+nums)
// 2、存在问题(bool)：dp[i]=dp[i] || dp[i-num]
// 3、组合问题：dp[i]+=dp[i-num]

// --------------------

// 动态规划 -- dp[i]表示i元的最少硬币数 -- dp[i] = min(dp[j]+dp[i-j])
func coinChange(coins []int, amount int) int {
	if coins == nil || len(coins) == 0 || amount < 0 {
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

	for _, coin := range coins { // 遍历每一个物品
		for i := 1; i <= amount; i++ { // 求在这个物品和之前的物品下的最优解
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
