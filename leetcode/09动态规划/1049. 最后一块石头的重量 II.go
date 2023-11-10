package main

// 1049. 最后一块石头的重量 II

// 有一堆石头，用整数数组stones 表示。其中stones[i] 表示第 i 块石头的重量。
//
// 每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为x 和y，且x <= y。
// 那么粉碎的可能结果如下：
//
// 如果x == y，那么两块石头都会被完全粉碎；
// 如果x != y，那么重量为x的石头将会完全粉碎，而重量为y的石头新重量为y-x。
// 最后，最多只会剩下一块 石头。返回此石头 最小的可能重量 。如果没有石头剩下，就返回 0。

// lastStoneWeightII .
// 动态规划 -- 01背包 -- 同leetcode 416
// 其实就是尽量让石头分成重量相同的两堆，相撞之后剩下的石头最小，这样就化解成01背包问题了，也就转化成416题了
//
// dp[i]表示容量为i的背包所能盛放的最大重量 -- 从选择列表中选择数字使其和尽量能够到总和的一半
// dp[i] = max(dp[i], dp[i - num] + num)
//
// 最后dp[target]里是容量为target的背包所能背的最大重量。
// 那么分成两堆石头，一堆石头的总重量是dp[target]，另一堆就是sum-dp[target]，
// 两堆石头相撞后剩余的石头重量就是 (sum-dp[target])-dp[target]
func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, num := range stones {
		sum += num
	}

	// 背包容量 -- 目标和向下取整，接下来计算容量为target的背包所能背的最大重量
	target := sum / 2

	// 构建dp数组
	dp := make([]int, target+1)

	// 初始化dp数组 -- 容量为0的背包所能盛放的最大重量为0
	dp[0] = 0

	// 开始遍历选择列表，更新每个容量的背包所能盛放的最大重量
	for _, num := range stones { // 遍历选择列表 -- 不断选择数字
		for i := target; i >= num; i-- { // 倒序遍历背包容量，计算并更新该容量的背包所能盛放的最大重量 -- 因为每个石头只能使用一次，所以需要倒序变量背包容量
			dp[i] = max(dp[i], dp[i-num]+num)
		}
	}

	// 最后dp[target]里是容量为target的背包所能背的最大重量。
	// 那么分成两堆石头，一堆石头的总重量是dp[target]，另一堆就是sum-dp[target]，
	// 两堆石头相撞后剩余的石头重量就是 (sum-dp[target])-dp[target]
	return (sum - dp[target]) - dp[target]
}
