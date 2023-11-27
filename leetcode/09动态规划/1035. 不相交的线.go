package main

// 1035. 不相交的线

// 在两条独立的水平线上按给定的顺序写下 nums1 和 nums2 中的整数。
// 现在，可以绘制一些连接两个数字 nums1[i]和 nums2[j]的直线，这些直线需要同时满足满足：
// nums1[i] == nums2[j]
// 且绘制的直线不与任何其他连线（非水平线）相交。
// 请注意，连线即使在端点也不能相交：每个数字只能属于一条连线。
//
// 以这种方法绘制线条，并返回可以绘制的最大连线数。

// maxUncrossedLines .
// 同 leetcode 1143
// 其实就是求两个数组的最长公共子序列的长度, 因为公共子序列数字的相对位置不变, 这样画出的线一定不相交
// 本题解法完全等同与1143，可以使用二维dp也可以使用一维dp，此处使用二维dp进行解答，一维dp解法详见1143
//
// dp[i][j]表示以下标i-1结尾的nums1和下标j-1结尾的nums2的子串的最长公共子序列长度
// dp[i][j] =
// 若 nums1[i-1] == nums2[j-1], 则 dp[i][j] = dp[i-1][j-1] + 1
// 若 nums1[i-1] != nums2[j-1], 则 dp[i][j] = max(dp[i-1][j], dp[i][j-1]) --> 公共子序列可以不连续, 因此不用像718那样使用0
func maxUncrossedLines(nums1 []int, nums2 []int) int {

	// 构造dp数组 -- dp[i][j]表示以下标i-1结尾的nums1和下标j-1结尾的nums2的子串的最长公共子序列长度
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	// 初始化dp数组 -- dp[0][0]、dp[i][0]、dp[0][j]其实都是没有意义的，只是为了dp结果正确
	// dp[0][0]、dp[i][0]、dp[0][j]的默认值应该为0，这些位置在构造dp数组的时候默认就是0，因此不需要单独初始化赋值为0

	// 开始dp -- 计算以每个下标结束的两个子串的公共子序列长度
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// 返回最终两个完整数组的公共子序列长度
	return dp[len(nums1)][len(nums2)]
}
