package main

// 718. 最长重复子数组

func main() {

}

// 2. 滚动数组优化，一维dp，空间占用较小 -- 类似礼物的最大价值、零钱兑换
func findLength(nums1 []int, nums2 []int) int {
	if nums1 == nil || len(nums1) == 0 || nums2 == nil || len(nums2) == 0 {
		return 0
	}

	dp := make([]int, len(nums1)+1)
	dp[0] = 0
	maxLen := 0
	for i := 1; i <= len(nums2); i++ {
		for j := len(nums1); j > 0; j-- {
			if nums1[j-1] == nums2[i-1] {
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
