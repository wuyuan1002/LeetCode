package main

// 673. 最长递增子序列的个数

// 给定一个未排序的整数数组，找到最长递增子序列的个数。

// func main() {

// }

// 类似 300
// 动态规划
func findNumberOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	} else if len(nums) <= 1 {
		return 1
	}

	dp := make([]int, len(nums))  // dp[i]表示下标为i的数结尾的最长上升序列长度
	ndp := make([]int, len(nums)) // ndp[i]表示下标为i的数结尾的最长上升序列个数

	for i := 0; i < len(nums); i++ {
		dp[i], ndp[i] = 1, 1
	}

	// 计算最长子序列长度和个数
	maxLen := 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					ndp[i] = ndp[j]
				} else if dp[j]+1 == dp[i] {
					ndp[i] += ndp[j]
				}
			}
			if dp[i] > maxLen {
				maxLen = dp[i]
			}
		}
	}

	// 统计最长子序列个数
	count := 0
	for i := 0; i < len(nums); i++ {
		if dp[i] == maxLen {
			count += ndp[i]
		}
	}

	return count
}
