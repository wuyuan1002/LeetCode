package main

// 53. 最大子数组和

// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
// 子数组 是数组中的一个连续部分。

// maxSubArray .
// 1. 动态规划
//
// dp[i]表示以下标i处数字结尾的数组的最大和
// dp[i] = max(dp[i-1]+nums[i], nums[i])
func maxSubArray(nums []int) int {
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

// maxSubArray1 .
// 2. 只使用一个变量currentSum保存前一个数字的最大和，遍历数组时计算最大子数组和
//
// 可以看到在计算下一个数字结尾的数组最大和时，只需要使用dp[i-1]即可，所以没必要使用dp数组记录所有中间结果
func maxSubArray1(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	currentSum, maxSum := nums[0], nums[0] // 以当前数字结尾的最大子数组和、全局最大子数组和
	for i := 1; i < len(nums); i++ {
		// 更新以当前数字结尾的最大子数组和
		currentSum = max(currentSum+nums[i], nums[i])
		// 更新全局最大子数组和
		maxSum = max(currentSum, maxSum)
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
