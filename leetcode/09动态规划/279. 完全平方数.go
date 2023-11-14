package main

import "math"

// 279. 完全平方数

// 给定正整数n，找到若干个完全平方数（比如1, 4, 9, 16, ...）使得它们的和等于 n。
// 你需要让组成和的完全平方数的个数最少。
//
// 给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。
// 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。
// 例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。

// numSquares .
// 动态规划 -- 完全背包
// 同leetcode 322, 只是物品不是零钱而是1、4、9这样的完全平方数了
// dp[i]表示组成i最少需要多少个完全平方数
// dp[i] = min(dp[i], dp[i - num*num] + 1)
func numSquares(n int) int {
	// 构造dp数组 -- 初始值使用 math.MaxInt32 表示每个中间数字都无法被组成
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}

	// 初始化dp数组 -- 组成0所需的完全平方数的个数是0
	dp[0] = 0

	// 开始dp -- 求组成每个数字最少需要多少个完全平方数
	for num := 1; num*num <= n; num++ { // 遍历选择列表 -- 小于n的所有完全平方数就是选择列表的所有数字，即遍历小于n的所有完全平方数(即num*num)
		for i := 1; i <= n; i++ { // 正序遍历背包容量 -- 因为每个完全平方数都可以被多次使用
			if i >= num*num {
				dp[i] = min(dp[i], dp[i-num*num]+1)
			}
		}
	}

	// 返回组成n所需的最少完全平方数个数
	if dp[n] == math.MaxInt32 {
		return 0
	}
	return dp[n]
}
