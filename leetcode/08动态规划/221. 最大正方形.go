package main

// 221. 最大正方形

// 在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。

// maximalSquare .
// dp[i][j]表示以matrix[i][j]作为右下角的最大正方形边长 -- 某一位置的最大正方形边长 = 它上、左、左上三个位置中边长的最小值 + 1
// dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
func maximalSquare(matrix [][]byte) int {

	// 构造dp数组 -- dp[i][j]表示以matrix[i][j]作为右下角的最大正方形边长
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}

	// 初始化dp数组 -- 第一行和第一列的坐标作为右下角时，其所能构成的最大正方形边长最大为1（就是其自身）
	// 此初始化过程可以在dp时一起做，所以不用单独循环进行初始化操作

	// 开始dp -- 求以每个坐标作为右下角的最大正方形边长，同时更新最大值
	maxSide := 0 // 最大边长
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			// 若当前位置不是'1'则直接跳过 -- 只有'1'的坐标才能构成正方形
			if matrix[i][j] != '1' {
				continue
			}

			// 若当前坐标为第一行或第一列 -- 其作为正方形右下角能构成的最大正方形边长就是1
			// 若不是第一行或第一列 -- 则其作为右下角能构成的最大边长 = 它上、左、左上三个位置中边长的最小值 + 1
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
			}

			// 更新最大边长
			maxSide = max(maxSide, dp[i][j])
		}
	}

	// 返回最大正方形面积
	return maxSide * maxSide
}
