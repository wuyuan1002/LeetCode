package main

import "fmt"

// 53. 最大子序和

// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

// 1. maxSubArray .
func maxSubArray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	currentSum, maxSum := nums[0], nums[0] // 当前和、最大和
	for i := 1; i < len(nums); i++ {
		currentSum = max(currentSum+nums[i], nums[i])
		maxSum = max(currentSum, maxSum)
	}
	return maxSum
}

// 2. maxSubArray1 动态规划
func maxSubArray1(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	maxSum := nums[0]            // 最大和
	dp := make([]int, len(nums)) // dp[i]表示以第i个数结尾的最大和
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		maxSum = max(dp[i], maxSum)
	}

	return maxSum
}
