package main

// 377. 组合总和 Ⅳ

// 给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。
// 请你从 nums 中找出并返回总和为 target 的元素组合的个数。
//
// 题目数据保证答案符合 32 位整数范围。

// combinationSum4 .
// 1. 回溯法  -- 可以求出每个解是什么-- leetcode 39
// 2. 动态规划 -- 完全背包(数字可以被多次使用)  -- 只能求出解的个数 -- leetcode 518
// 组合问题：在选择列表中选择元素，求和为target的解的个数
// dp[i]表示和为i的排列个数
// dp[i] = dp[i] + dp[i - num]
func combinationSum4(nums []int, target int) int {
	// 构造dp数组 -- dp[i]表示和为i的排列个数
	dp := make([]int, target+1)

	// dp数组初始化 -- 此处dp[0]无意义，只是为了后续dp结果正确
	dp[0] = 1

	// 开始dp -- 计算各种背包容量的排列个数
	// 本题需要先遍历背包容量后遍历选择列表 -- 因为本题求的是排列问题，需要考虑解的顺序，而leetcode 518 不需要考虑解的顺序
	for i := 1; i <= target; i++ { // 正序遍历背包容量累加计算构成每个容量的排列个数 -- 若每个数字只能使用一次则需要倒序遍历背包容量
		for _, num := range nums { // 遍历选择列表 -- 选择每一个元素
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}

	// 返回和为target的排列个数
	return dp[target]
}
