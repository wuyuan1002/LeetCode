package main

// 674. 最长连续递增序列

// 给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。
//
// 连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，
// 都有 nums[i] < nums[i + 1] ，那么子序列 [nums[l], nums[l + 1], ...,
// nums[r - 1], nums[r]] 就是连续递增子序列。

// findLengthOfLCIS .
// dp[i]表示以第i个数字结尾的最长连续递增序列的长度
// dp[i] = dp[i-1] + 1
func findLengthOfLCIS(nums []int) int {
	// 总结果 -- 最长连续递增子序列长度 -- 因为每个数字本身长度就是1，所以总结果最小也是1
	maxLen := 1

	// 构造dp数组 -- dp[i]表示以第i个数字结尾的最长连续递增序列的长度
	dp := make([]int, len(nums))

	// 初始化dp数组 -- 以每个数字结尾的最长递增长度为1（数字本身长度为1）
	for i := range dp {
		dp[i] = 1
	}

	// 开始dp -- 求以每个数字结尾的最长连续递增子序列长度
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			// 若当前数字比前一个数字大，则以当前数字结尾的最长连续递增子序列长度为前一个数字的长度+1
			dp[i] = dp[i-1] + 1
		}
		// 更新整体的最长连续递增子序列长度
		maxLen = max(maxLen, dp[i])
	}

	return maxLen
}
