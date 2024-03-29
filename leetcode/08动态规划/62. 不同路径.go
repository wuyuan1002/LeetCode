package main

// 62. 不同路径

// 一个机器人位于一个 m x n网格的左上角 （起始点在下图中标记为 “Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
// 问总共有多少条不同的路径？

// uniquePaths .
// 二维dp
// dp[i][j]表示在[i, j]位置的走法个数
// 到达某一个位置的走法 == 它上面位置的走法个数 + 左面位置的走法个数
// dp[i][j] = dp[i][j-1] + dp[i-1][j]
func uniquePaths(m int, n int) int {
	if m <= n || n <= 0 {
		return 0
	}

	// 构造dp数组
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 初始化dp数组，第一列和第一行的走法个数都为1
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	// 遍历计算每个位置的走法个数并保存中间结果到dp数组
	for i := 1; i < m; i++ { // 遍历每一行
		for j := 1; j < n; j++ { // 遍历当前行的每一列
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}

	// 返回最后一个位置的走法个数
	return dp[m-1][n-1]
}

// uniquePaths1 .
// 滚动数组 -- 一维dp
// 从上面的二维dp我们可以发现，每次遍历下一行时，其实只使用前一行的数据，对于前两行前三行等的数据其实已经没用了
// 我们可以使用一维数组滚动更新每一行的走法个数，计算下一行时直接使用当前数组数字计算后赋值即可
// dp[j]表示在某一行第j个位置的走法个数
// dp[j] = dp[j] + dp[j - 1]
func uniquePaths1(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	// 构造dp数组 -- dp数组记录的是前一行每个位置的走法个数
	dp := make([]int, n)

	// 初始化dp数组 -- 第一行第一个位置的走法个数为1
	dp[0] = 1

	// 遍历计算每个位置的走法个数并保存中间结果到dp数组
	for i := 0; i < m; i++ { // 遍历每一行
		for j := 0; j < n; j++ { // 遍历当前行的每一列
			if j == 0 {
				// 当前行的第一列 -- 始终等于前一行第一列的值，因为第一列只能从前一行向下移动过来
				dp[j] = dp[j]
			} else {
				// 当前行第j列的走法个数 == 前一行第j列的走法个数(dp[j]) + 当前行前一列的走法个数(dp[j-1])
				// 计算完当前行第j列的走法个数后，将结果更新到dp数组中，供下一行计算使用
				dp[j] = dp[j] + dp[j-1]
			}
		}
	}

	// 遍历完后，dp数组存的就是最后一行每一列的走法个数，返回最后一列的走法个数
	return dp[n-1]
}
