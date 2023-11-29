package main

// 115. 不同的子序列

// 给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。
// 字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。
// （例如，"ACE"是"ABCDE"的一个子序列，而"AEC"不是）
// 题目数据保证答案符合 32 位带符号整数范围。

// numDistinct .
// dp[i][j]表示s的前i-1个字符中，t的前j-1个字符出现的次数
// dp[i][j] =
// 若 s[i-1] == t[j-1], 则 dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
// 若 s[i-1] != t[j-1], 则 dp[i][j] = dp[i-1][j]
func numDistinct(s string, t string) int {

	// 构造dp数组 -- dp[i][j]表示s的前i-1个字符中，t的前j-1个字符出现的次数
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}

	// 初始化dp数组 -- dp[i][0]都为1，表示s的每个子串中空串的个数为1
	for i := 0; i <= len(s); i++ {
		dp[i][0] = 1
	}

	// 开始dp -- 分别求s的每个子串出现t的每个子串的个数
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	// 返回最终s中出现t的个数
	return dp[len(s)][len(t)]
}
