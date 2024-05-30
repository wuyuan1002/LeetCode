package main

// 583. 两个字符串的删除操作

// 给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，
// 每步可以删除任意一个字符串中的一个字符。

// minDistance .
// 1. 求两个字符串的最长公共子序列，然后两串长度与公共子序列差值的和就是最小的删除次数 -- 同 leetcode 1143
// 2. 编辑距离 -- 同 leetcode 72
func minDistance583(word1 string, word2 string) int {

	// 构造dp数组 -- dp[i][j]表示以下标i-1结尾的word1和下标j-1结尾的word2的子串的最长公共子序列长度
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	// 初始化dp数组 -- dp[0][0]、dp[i][0]、dp[0][j]其实都是没有意义的，只是为了dp结果正确
	// dp[0][0]、dp[i][0]、dp[0][j]的默认值应该为0，这些位置在构造dp数组的时候默认就是0，因此不需要单独初始化赋值为0

	// 开始dp -- 计算以每个下标结束的两个子串的公共子序列长度
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// 两串长度与公共子序列差值的和就是最小的删除次数
	maxLen := dp[len(word1)][len(word2)]
	return len(word1) - maxLen + len(word2) - maxLen
}
