package main

// 718. 最长重复子数组

// 给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。
//
// 输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
// 输出：3
// 解释：长度最长的公共子数组是 [3,2,1] 。

// findLength .
// 二维dp
//
// dp[i][j]表示以下标i-1结尾的nums1和下标j-1结尾的nums2的数组的重复子数组长度
// dp[i][j] =
// 若 nums1[i-1] == nums2[j-1], 则 dp[i][j] = dp[i-1][j-1] + 1
// 若 nums1[i-1] != nums2[j-1], 则 dp[i][j] = 0 --> 重复子数组需要连续, 因此不能像1143那样使用max(dp[i-1][j], dp[i][j-1])
func findLength(nums1 []int, nums2 []int) int {

	// 最长重复子数字长度
	maxLen := 0

	// 构造dp数组 -- dp[i][j]表示以下标i-1结尾的nums1和下标j-1结尾的nums2的数组的重复子数组长度
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	// 初始化dp数组 -- dp[0][0]、dp[i][0]、dp[0][j]其实都是没有意义的，只是为了dp结果正确
	// dp[0][0]、dp[i][0]、dp[0][j]的默认值应该为0，这些位置在构造dp数组的时候默认就是0，因此不需要单独初始化赋值为0

	// 开始dp -- 求以每个下标结束的两个子数组的公共子数组长度，同时更新全局最长公共子数组长度
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				// 若当前数字相等，则以当前下标结尾的子数组公共长度为前一个子数组公共长度+1
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 若当前值不相等，则以当前下标结尾的子数组公共长度为0
				dp[i][j] = 0
			}

			// 更新全局最长重复子数组长度
			maxLen = max(maxLen, dp[i][j])
		}
	}

	return maxLen
}

// findLength1 .
// 一维dp
// 可以看出dp[i][j]（当前行的结果）都是由dp[i-1][j-1]（左上方的结果）得出，也就是下一行的结果只依赖上一行的结果，
// 因此可以只使用一维数组记录上一行的结果即可，滚动数组，也就是相当于可以把上一层dp[i-1][j]拷贝到下一层dp[i][j]来继续用，
// 因为要使用的是上一行的左上方的结果，那么遍历nums2数组的时候要从后向前遍历，避免重复覆盖前一行左上方的结果
//
// dp[j]表示以下标j-1结尾的nums2子数组和以下标i-1结尾的nums1子数组的重复子数组长度
// dp[j] =
// 若 nums1[i-1] == nums2[j-1], 则 dp[j] = dp[j - 1] + 1
// 若 nums1[i-1] != nums2[j-1], 则 dp[j] = 0
func findLength1(nums1 []int, nums2 []int) int {

	// 最长重复子数字长度
	maxLen := 0

	// 构造dp数组 -- dp[j]表示以下标i-1结尾的nums1子数组和以下标j-1结尾的nums2子数组的重复子数组长度
	dp := make([]int, len(nums2)+1)

	// 初始化dp数组 -- 初始时保持dp数组的每个位置都为0、dp[0]其实是没有意义的，只是为了dp结果正确

	// 开始dp -- 求以下标i-1结尾的nums1子数组和以下标j-1结尾的nums2子数组的重复子数组长度
	for i := 1; i <= len(nums1); i++ {
		// 求nums2[j-1]和nums1[i-1]的重复子数组长度 -- 倒序遍历nums2，防止覆盖前一行的结果
		for j := len(nums2); j >= 1; j-- {
			if nums1[i-1] == nums2[j-1] {
				dp[j] = dp[j-1] + 1
			} else {
				dp[j] = 0
			}

			// 更新全局最长重复子数组长度
			maxLen = max(maxLen, dp[j])
		}
	}

	return maxLen
}
