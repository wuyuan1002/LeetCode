package main

// 53. 最大子数组和

// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
// 子数组 是数组中的一个连续部分。

// maxSubArray .
// 1. 使用两个变量分别记录当前和以及最大和，在遍历数组过程中更新两个值
// 2. 动态规划
func maxSubArray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	currentSum, maxSum := nums[0], nums[0] // 当前的最大连续和、最大子数组和
	for i := 1; i < len(nums); i++ {
		// 更新当前最大的连续和
		currentSum = max(currentSum+nums[i], nums[i])
		// 更新全局最大子数组和
		maxSum = max(currentSum, maxSum)
	}

	return maxSum
}

// maxSubArrayDP 动态规划
func maxSubArrayDP(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	maxSum := nums[0] // 最大子数组和

	// 定义并初始化dp数组 -- dp[i]表示以第i个数结尾的最大子数组和
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	// 遍历数组，更新dp数组和最大子数组和
	for i := 1; i < len(nums); i++ {
		// 更新以当前数字结尾的最大子数组和
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		// 更新全局的最大子数组和
		maxSum = max(dp[i], maxSum)
	}

	return maxSum
}

// max .
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
