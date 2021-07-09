package main

// 718. 最长重复子数组

// 给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。

func main() {

}

// 1. 动态规划
// 2。 滑动窗口

// 2. 二维dp，空间占用较大
func findLength(nums1 []int, nums2 []int) int {
	if nums1 == nil || len(nums1) == 0 || nums2 == nil || len(nums2) == 0 {
		return 0
	}
	// dp[i][j] =
	// 若 A[i-1] != B[j-1], 则 dp[i][j] = 0
	// 若 A[i-1] == B[j-1], 则 dp[i][j] = dp[i-1][j-1] + 1

	maxLen := 0

	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}
	dp[0][0] = 0

	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}

			if dp[i][j] > maxLen {
				maxLen = dp[i][j]
			}
		}
	}

	return maxLen
}

// 2. 滚动数组优化，一维dp，空间占用较小 -- 类似礼物的最大价值、零钱兑换
func findLength1(nums1 []int, nums2 []int) int {
	if nums1 == nil || len(nums1) == 0 || nums2 == nil || len(nums2) == 0 {
		return 0
	}

	maxLen := 0

	dp := make([]int, len(nums2)+1)
	dp[0] = 0

	for i := 1; i <= len(nums1); i++ {
		for j := len(nums2); j >= 1; j-- {
			if nums1[i-1] == nums2[j-1] {
				dp[j] = dp[j-1] + 1
			} else {
				dp[j] = 0
			}

			if dp[j] > maxLen {
				maxLen = dp[j]
			}
		}
	}

	return maxLen
}
