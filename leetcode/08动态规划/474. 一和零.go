package main

import "strings"

// 474. 一和零

// 给你一个二进制字符串数组 strs 和两个整数 m 和 n 。
// 请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 。
//
// 如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。

// findMaxForm .
// 动态规划 01背包问题，一般的动态规划为给定一个选择列表和一个target约束，求解达到目标的最优解，
// 而此题的差异是给了一个选择列表和两个target约束，因此在dp时要比普通动态规划多一个纬度
//
// dp[i][j][k]表示在前i个字符串中，使用j个0和k个1的情况下最多可以得到的字符串数量 -- 最终答案为dp[len(strs)][m][n]
// dp[i][j][k] = max(dp[i−1][j][k], dp[i−1][j−zeros][k−ones] + 1)
func findMaxForm(strs []string, m int, n int) int {
	// 构造一个dp数组，用来保存中间计算结果
	dp := make([][][]int, len(strs)+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}

	// 开始dp，从选择列表中不断选择字符串
	for i, s := range strs {
		// 计算当前字符串中0和1的个数
		count0 := strings.Count(s, "0")
		count1 := len(s) - count0

		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				dp[i+1][j][k] = dp[i][j][k]
				if j >= count0 && k >= count1 {
					dp[i+1][j][k] = max(dp[i+1][j][k], dp[i][j-count0][k-count1]+1)
				}
			}
		}
	}

	return dp[len(strs)][m][n]
}
