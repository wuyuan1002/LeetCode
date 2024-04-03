package main

// 931. 下降路径最小和

// 给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。
//
// 下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素
// 和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第一个元素）。具体来说，
// 位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1, col + 1) 。

// minFallingPathSum .
// dp[i][j]表示从第一行下降到matrix[i][j]位置的最小和
// dp[i][j] = min(dp[i-1][j], min(dp[i-1][j-1], dp[i-1][j+1])) + matrix[i][j]
//
// 由于当前行dp[i][j]的计算只依赖上一行的结果，所以可以使用滚动数组只保留前一行的结果即可
func minFallingPathSum(matrix [][]int) int {
	// 构造dp数组 -- dp[i][j]表示从第一行下降到matrix[i][j]位置的最小和
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}

	// 初始化dp数组 -- 第一行的下降路径和本身为自身的值
	for j := 0; j < len(matrix[0]); j++ {
		dp[0][j] = matrix[0][j]
	}

	// 从第一行开始dp，计算每一个位置的下降路径
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if j == 0 { // 第一列只能依赖 上、右上
				dp[i][j] = min(dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
			} else if j == len(matrix[0])-1 { // 最后一列只能依赖 左上、上
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + matrix[i][j]
			} else { // 其他列依赖 左上、上、右上
				dp[i][j] = min(dp[i-1][j], min(dp[i-1][j-1], dp[i-1][j+1])) + matrix[i][j]
			}
		}
	}

	// 返回最后一行的最小下降路径和
	result := dp[len(matrix)-1][0]
	for _, sum := range dp[len(matrix)-1] {
		result = min(result, sum)
	}
	return result
}
