package main

// 583. 两个字符串的删除操作

// 给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，
// 每步可以删除任意一个字符串中的一个字符。

// func main() {

// }

// 1. 类似 1143.最长公共子序列
// 2. 编辑距离
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	// 求两串的最长公共子序列
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// 两串长度与子序列长度差值的大的那个就是最小的删除次数
	maxLen := dp[len(word1)][len(word2)]
	return len(word1) - maxLen + len(word2) - maxLen
}
