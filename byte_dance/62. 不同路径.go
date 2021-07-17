package main

// 62. 不同路径

// 一个机器人位于一个 m x n网格的左上角 （起始点在下图中标记为 “Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//
// 问总共有多少条不同的路径？

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
			if j == 0 { // 第一列
				dp[j] = dp[j]
			} else { // 其他
				dp[j] = dp[j-1] + dp[j]
			}
		}
	}

	return dp[n-1]
}
