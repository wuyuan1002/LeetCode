package main

// 53. 最大子数组和

// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
// 子数组 是数组中的一个连续部分。

// maxSubArray .
// 1. 使用两个变量分别记录当前和以及最大和，在遍历数组过程中更新两个值
func maxSubArray(nums []int) int {
	// 以当前数字结尾的最大子数组和、全局最大子数组和
	currentSum, maxSum := nums[0], nums[0]

	// 遍历每个数字，更新以当前数字结尾的最大子数组和，并更新全局最大子数组和
	for i := 1; i < len(nums); i++ {
		// 更新以当前数字结尾的最大子数组和
		currentSum = max(currentSum+nums[i], nums[i])
		// 更新全局最大子数组和
		maxSum = max(currentSum, maxSum)
	}

	return maxSum
}

// maxSubArrayDP .
// 2. 动态规划
// dp[i]表示以第i个数字结尾的最大子数组和
// dp[i] = max(dp[i-1]+nums[i], nums[i])
func maxSubArrayDP(nums []int) int {
	// 全局最大子数组和
	maxSum := nums[0]

	// 定义dp数组 -- dp[i]表示以第i个数字结尾的最大子数组和
	dp := make([]int, len(nums))

	// 初始化dp数组 -- 以第一个数字结尾的最大子数组和就是该数字本身
	dp[0] = nums[0]

	// 开始dp -- 更新dp数组中以每个数字结尾的最大子数组和，并更新全局最大子数组和
	for i := 1; i < len(nums); i++ {
		// 更新以当前数字结尾的最大子数组和
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		// 更新全局的最大子数组和
		maxSum = max(dp[i], maxSum)
	}

	return maxSum
}
