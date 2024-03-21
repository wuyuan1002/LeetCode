package main

// 746. 使用最小花费爬楼梯

// 给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。
// 一旦你支付此费用，即可选择向上爬一个或者两个台阶。
//
// 你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。
//
// 请你计算并返回达到楼梯顶部的最低花费。

// minCostClimbingStairs .
// 可以选择从下标为0或下标为1的台阶开始爬楼梯，也就是在第0和第1个台阶不需要费用
// cost[i]表示从第i个台阶向上爬所需的费用
// dp[i]表示到达第i个台阶的最小花费 -- 可以是从前一个台阶跳上来也可以是从前两个台阶跳上来
// dp[i] = min(dp[i - 1] + cost[i - 1], dp[i - 2] + cost[i - 2])
func minCostClimbingStairs(cost []int) int {
	// 构造dp数组 -- dp[i]表示到达第i个台阶的最小花费
	dp := make([]int, len(cost)+1)

	// 初始化dp数组 -- 因为到达第0和第1个台阶不需要费用，所以在前两个的花费为0
	dp[0] = 0
	dp[1] = 0

	// 遍历cost数组，挨个计算到达每个台阶的最小花费
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	// 返回到达最后一个台阶的花费
	return dp[len(cost)]
}

// min .
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
