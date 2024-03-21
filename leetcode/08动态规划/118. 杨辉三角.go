package main

// 118. 杨辉三角

// 给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
//
// 在「杨辉三角」中，每个数是它左上方和右上方的数的和。

// generate .
func generate(numRows int) [][]int {
	// 创建dp数组
	dp := make([][]int, numRows)

	// 遍历每一行，结算并记录该行结果
	for i := 0; i < numRows; i++ {
		// 创建第i行的数组用于存结果
		dp[i] = make([]int, i+1)

		// 遍历第i行，计算每个位置的结果
		for j := 0; j <= i; j++ {
			if j == 0 || j == i { // 第一列和最后一列
				dp[i][j] = 1
			} else { // 中间列
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			}
		}
	}

	return dp
}
