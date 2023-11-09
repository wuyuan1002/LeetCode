package main

// 518. 零钱兑换 II

// 给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
// 请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。
//
// 假设每一种面额的硬币有无限个。
// 题目数据保证结果符合 32 位带符号整数。

// change .
// 因为只需要求出解的个数，而不是求出每个解是什么，所以不必使用回溯法解决 -- leetcode 39
// 完全背包 -- 因为每种硬币数量是无限的，说明每个面额的硬币可以被多次选择
// dp[i]表示凑出总金额i的硬币组合数
// dp[i] = dp[i] + dp[i - coin]
func change(amount int, coins []int) int {
	// 构造dp数组 -- dp[i]表示凑出总金额i有几种组合方式
	dp := make([]int, amount+1)
	// 初始化dp数组 -- 凑出总金额为0的硬币组合数为1，其实这个值没有意义，只是如果dp[0] = 0 的话，后面所有推导出来的值都是0了
	dp[0] = 1

	// 开始遍历选择列表和背包容量，计算每个容量的组合个数
	for _, coin := range coins { // 遍历选择列表
		for i := 1; i <= amount; i++ { // 遍历背包容量 -- 若每个元素只能选择一次（01背包）则背包容量应该倒序遍历
			if i >= coin {
				dp[i] += dp[i-coin]
			}
		}
	}

	// 返回凑出总金额为amount时的硬币组合数
	return dp[amount]
}
