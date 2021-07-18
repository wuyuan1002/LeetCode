package main

import (
	"fmt"
)

// 1143. 最长公共子序列

// 给定两个字符串text1 和text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
//
// 一个字符串的子序列是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
//
// 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
// 两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。

func main() {
	fmt.Println(longestCommonSubsequence1("abcba", "abcbcba"))
}

// 类似于 718. 最长重复子数组
// 动态规划，二维dp -- dp[i][j]表示以下标i-1结尾的text1和下标j-1结尾的text2的数组的字符串的公共子序列长度
func longestCommonSubsequence(text1 string, text2 string) int {
	if text1 == "" || text2 == "" {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	// dp[i][j] =
	// 若 A[i-1] == B[j-1], 则 dp[i][j] = dp[i-1][j-1] + 1
	// 若 A[i-1] != B[j-1], 则 dp[i][j] = max(dp[i-1][j], dp[i][j-1])

	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	dp[0][0] = 0

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}

// 滚动数组，一维dp
func longestCommonSubsequence1(text1 string, text2 string) int {
	if text1 == "" || text2 == "" {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	dp := make([]int, len(text2)+1)
	dp[0] = 0

	for i := 1; i <= len(text1); i++ {
		upLeft := 0
		for j := 1; j <= len(text2); j++ {
			tmp := dp[j]
			if text1[i-1] == text2[j-1] {
				dp[j] = upLeft + 1
			} else {
				dp[j] = max(dp[j-1], dp[j])
			}
			upLeft = tmp // 记住左上角的值，防止被覆盖
		}
	}

	return dp[len(dp)-1]
}
