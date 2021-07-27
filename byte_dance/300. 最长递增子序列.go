package main

// 300. 最长递增子序列

// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
// 子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。
// 例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

func main() {

}

// 类似 Hot100 128
// 动态规划 -- dp[i]表示以第i个数字结尾的最长递增序列的长度 -- dp[i] = max(dp[i], dp[j] + 1) for j in [0, i)
func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	dp := make([]int, len(nums)) // dp[i]表示以第i个数字结尾的最长递增序列长度, dp[i] = max(dp[i], dp[j] + 1)
	for i := range dp {
		dp[i] = 1
	}

	maxLen := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])
	}

	return maxLen
}
