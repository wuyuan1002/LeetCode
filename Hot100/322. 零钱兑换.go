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

// 背包问题: 给定一个target(直接或间接求出)，给定一个选择列表，求达到target的最优解、解的个数、是否存在等问题

// 背包问题的转移方程:
// 1、最值问题(求最优解): dp[i] = max/min(dp[i], dp[i-nums]+1) 或 dp[i] = max/min(dp[i], dp[i-num]+nums)
// Hot100 322

// 2、存在问题(是否存在): dp[i]=dp[i] || dp[i-num]
// Hot100 416

// 3、组合问题(解的个数): dp[i]+=dp[i-num]
// 只是求出解的个数，而不是求出每个解是什么 -- 见Hot100 39
// Hot100 518、39、494

// --------------------

// 类似Hot100 518, 416，39(与回溯法的区别)
// 动态规划 -- dp[i]表示i元的最少硬币数 -- dp[i] = min(dp[i], dp[i-num]+1)
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
		for i := 1; i <= amount; i++ { // 求选择这个物品后构成目标值的最优解 -- 每个物品只能被使用1次时应该是倒序遍历，如Hot100 416、494
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
