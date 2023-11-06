package main

import "math"

// 322. 零钱兑换

// 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
//
// 你可以认为每种硬币的数量是无限的。

// coinChange .
// 完全背包 -- 因为每种硬币数量是无限的，说明每个面额的硬币可以被多次选择
// dp[i]表示凑足i元需要的最少硬币数
// dp[i] = min(dp[i], dp[i-num]+1)
func coinChange(coins []int, amount int) int {
	// 构造dp数组 -- 初始值使用 math.MaxInt32 表示每种总金额都无法被组成
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}

	// 初始化dp数组 -- 凑足总金额为0所需钱币个数是0
	dp[0] = 0

	for _, coin := range coins { // 遍历选择列表中的每一种物品
		for i := 1; i <= amount; i++ { // 遍历背包容量 -- 求选择这个物品后构成目标值的最优解 -- 每个物品只能被使用1次时应该是倒序遍历，如leetcode 416、494
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	// 返回凑足amount需要的最少硬币数
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
