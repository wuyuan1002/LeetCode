package main

// 1143. 最长公共子序列

// 给定两个字符串text1 和text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
//
// 一个字符串的子序列是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
//
// 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
// 两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。

// longestCommonSubsequence .
// 同 leetcode 1035
// 二维dp
//
// dp[i][j]表示以下标i-1结尾的text1和下标j-1结尾的text2的子串的最长公共子序列长度
// dp[i][j] =
// 若 text1[i-1] == text2[j-1], 则 dp[i][j] = dp[i-1][j-1] + 1
// 若 text1[i-1] != text2[j-1], 则 dp[i][j] = max(dp[i-1][j], dp[i][j-1]) --> 公共子序列可以不连续, 因此不用像718那样使用0
func longestCommonSubsequence(text1 string, text2 string) int {

	// 构造dp数组 -- dp[i][j]表示以下标i-1结尾的text1和下标j-1结尾的text2的子串的最长公共子序列长度
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	// 初始化dp数组 -- dp[0][0]、dp[i][0]、dp[0][j]其实都是没有意义的，只是为了dp结果正确
	// dp[0][0]、dp[i][0]、dp[0][j]的默认值应该为0，这些位置在构造dp数组的时候默认就是0，因此不需要单独初始化赋值为0

	// 开始dp -- 计算以每个下标结束的两个子串的公共子序列长度
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// 返回最终两个完整字符串的公共子序列长度
	return dp[len(text1)][len(text2)]
}

// longestCommonSubsequence1 .
// 一维dp
// 可以看出dp[i][j]（当前行的结果）都是由dp[i-1][j-1]（左上方的结果）得出，也就是下一行的结果只依赖上一行的结果，
// 因此可以只使用一维数组记录上一行的结果即可，滚动数组，也就是相当于可以把上一层dp[i-1][j]拷贝到下一层dp[i][j]来继续用，
// 因为要使用的是上一行的左上方的结果，那么遍历text2的时候需要使用一个变量记录一下左上方的结果，防止被覆盖
//
// 为什么不能倒序遍历text2，而是使用一个变量记录左上角的值呢？
// 因为在两个字符不想等时需要使用上面和左面的值，而不是像718那样直接为0，如果使用倒序遍历的话，前面的值始终是0，无法使用前面的值
//
// dp[j]表示以下标j-1结尾的text2和下标i-1结尾的text1的子串的最长公共子序列长度
// dp[j] =
// 若 nums1[i-1] == nums2[j-1], 则 dp[j] = dp[j - 1] + 1
// 若 nums1[i-1] != nums2[j-1], 则 dp[j] = max(dp[j-1], dp[j])
func longestCommonSubsequence1(text1 string, text2 string) int {
	// 构造dp数组 -- dp[j]表示以下标j-1结尾的text2和下标i-1结尾的text1的子串的公共子序列长度
	dp := make([]int, len(text2)+1)

	// 初始化dp数组 -- 初始时保持dp数组的每个位置都为0、dp[0]其实是没有意义的，只是为了dp结果正确

	// 开始dp -- 求以下标i-1结尾的nums1子数组和以下标j-1结尾的nums2子数组的重复子数组长度
	for i := 1; i <= len(text1); i++ {
		// 用于记录当前行的左上角的值, 防止被覆盖 -- 第一列的左上角的值始终为0（类似于dp[i][0]的值始终为0）
		upLeft := 0

		// 求text2[j-1]和text1[i-1]的最长公共子序列长度
		for j := 1; j <= len(text2); j++ {
			// 记住前一行当前列的值（这个值就是当前行下一列左上角的值），因为后面dp[j]会被覆盖成当前行当前列的值，所以先使用临时变量记录一下
			temp := dp[j]

			if text1[i-1] == text2[j-1] {
				dp[j] = upLeft + 1
			} else {
				dp[j] = max(dp[j-1], temp)
			}

			// 更新下一列左上角的值为前一行当前列的值
			upLeft = temp
		}
	}

	// 返回最终两个完整字符串的公共子序列长度
	return dp[len(text2)]
}
