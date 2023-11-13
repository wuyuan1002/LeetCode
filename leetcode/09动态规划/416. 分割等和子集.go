package main

// 416. 分割等和子集

// 给你一个 只包含正整数 的 非空 数组 nums 。
// 请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

// canPartition .
// 存在问题：在选择列表中选择元素，是否存在和为target的解
//
// 在集合中找到一些数字使得它们的和为整体总和的一半，说明可以将数组分割为等和子集
// 01背包问题 -- 存在问题 -- dp[i] = dp[i] || dp[i-num]
// dp[i]表示是否存在子集和为i
// dp[i] = dp[i] || dp[i-num]
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		// 总和为奇数，无法分成相等的两部分
		return false
	}

	target := sum / 2 // 背包容量 -- 每个子数组的目标和 -- 零钱兑换里的总金额

	// 构造dp数组 -- dp[i]表示是否存在子集和为i
	dp := make([]bool, target+1)

	// 初始化dp数组 -- 目标和为0不需要选择任何元素，所以是可以实现的
	dp[0] = true

	// 开始dp -- 不断从选择列表中选择元素，判断是否可以构成目标值
	for _, num := range nums { // 遍历选择列表 -- 所有零钱
		for i := target; i >= num; i-- { // 遍历背包容量 -- 每个物品只能被使用1次时就应该倒序遍历，如leetcode 322
			dp[i] = dp[i] || dp[i-num]
		}
	}

	// 返回是否存在子集和为target
	return dp[target]
}
