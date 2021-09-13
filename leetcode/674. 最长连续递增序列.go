package main

// 674. 最长连续递增序列

// 给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。
//
// 连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，
// 都有 nums[i] < nums[i + 1] ，那么子序列 [nums[l], nums[l + 1], ...,
// nums[r - 1], nums[r]] 就是连续递增子序列。

func main() {

}

// 类似 128、300
// 动态规划 -- dp[i]表示以下标i处数字结尾的连续递增序列长度
func findLengthOfLCIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	maxLen := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		maxLen = max(maxLen, dp[i])
	}

	return maxLen
}

// 第二次做
func findLengthOfLCIS1(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	// 构建并初始化dp
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	// 遍历数组
	maxLen := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		maxLen = max(maxLen, dp[i])
	}

	return maxLen
}
