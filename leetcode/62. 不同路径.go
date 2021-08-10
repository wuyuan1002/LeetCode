package main

// 62. 不同路径

func main() {

}

// 类似于offer 47
// 动态规划 -- 到达某一个位置的走法 == 它上面位置的走法个数 + 左面位置的走法个数
func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	dp := make([]int, n)
	dp[0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j > 0 {
				dp[j] = dp[j] + dp[j-1]
			}
		}
	}

	return dp[n-1]
}
