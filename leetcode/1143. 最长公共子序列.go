package main

// 1143. 最长公共子序列

// func main() {

// }

// 二维dp
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text2)+1)
	for i := range dp {
		dp[i] = make([]int, len(text1)+1)
	}

	for i := 1; i <= len(text2); i++ {
		for j := 1; j <= len(text1); j++ {
			if text2[i-1] == text1[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[len(text2)][len(text1)]
}

// 滚动数组，一维dp
func longestCommonSubsequence1(text1 string, text2 string) int {
	dp := make([]int, len(text1)+1)
	dp[0] = 0

	for i := 1; i <= len(text2); i++ {
		upLeft := 0
		for j := 1; j <= len(text1); j++ {
			tmp := dp[j]
			if text2[i-1] == text1[j-1] {
				dp[j] = upLeft + 1
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			upLeft = tmp
		}
	}

	return dp[len(text1)]
}
