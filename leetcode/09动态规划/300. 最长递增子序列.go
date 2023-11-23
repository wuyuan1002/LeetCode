package main

// 300. 最长递增子序列

// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
// 子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。
// 例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

// lengthOfLIS .
// dp[i]表示以第i个数字结尾的最长递增序列的长度
// dp[i] = max(dp[i], dp[j] + 1)
func lengthOfLIS(nums []int) int {
	// 总结果 -- 最长递增子序列长度 -- 因为每个数字本身长度就是1，所以总结果最小也是1
	maxLen := 1

	// 构造dp数组 -- dp[i]表示以第i个数字结尾的最长递增序列的长度
	dp := make([]int, len(nums))

	// 初始化dp数组 -- 以每个数字结尾的最长递增长度为1（数字本身长度为1）
	for i := range dp {
		dp[i] = 1
	}

	// 开始dp -- 求以每个数字结尾的最长递增子序列长度，并在过程中计算整体最长的递增子序列长度
	for i := 1; i < len(nums); i++ {
		// 计算以第i个数字结尾的最长递增子序列长度 == [0, i)各个位置的最长递增子序列长度+1 的最大值
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])
	}

	return maxLen
}
