package main

// 53. 最大子序和

// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

func main() {

}

// 1.
func maxSubArray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	maxSum := nums[0]     // 最大和
	currentSum := nums[0] // 当前和
	for i := 1; i < len(nums); i++ {
		currentSum = max(currentSum+nums[i], nums[i])
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

// 2. 动态规划
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

	maxSum := nums[0] // 最大和

	dp := make([]int, len(nums)) // dp[i]表示以第i个数结尾的最大和
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		maxSum = max(maxSum, dp[i])
	}

	return maxSum
}
