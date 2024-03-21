package main

// 64. 最小路径和

// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
// 说明：每次只能向下或者向右移动一步。

// minPathSum .
// 当前位置的最小和取决于我的上面和左面位置的最小和
// dp[i][j]表示grid[i][j]位置的最小路径和
// dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
func minPathSum(grid [][]int) int {

	// 构造dp数组 -- dp[i][j]表示grid[i][j]位置的最小路径和
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}

	// 开始dp，逐个计算每个坐标的最小路径和
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 { // 第一个坐标
				dp[i][j] = grid[i][j]
			} else if j == 0 { // 第一列
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else if i == 0 { // 第一行
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else {
				dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// 返回最后一个坐标的最小路径和
	return dp[len(grid)-1][len(grid[0])-1]
}

// minPathSum1 .
// 一维dp
// 可以看出dp[i][j]（当前行的结果）都是由dp[i-1][j]、dp[i][j-1]（左面位置和上面位置的结果）得出，也就是下一行的结果只依赖当前行和上一行的结果，
// 因此可以只使用一维数组记录上一行的结果即可，滚动数组，也就是相当于可以把上一层dp[i-1][j]拷贝到下一层dp[i][j]来继续用，
//
// 因为每个坐标依赖上面和当前行左面的结果，所以需要顺序遍历每一行来更新当前行每个位置j的最小路径和，这样下一个坐标j+1才能使用当前行左面的数据，不然dp[j-1]还是上一行的数据
//
// dp[j]表示grid第i行各个位置j（即grid[i][j]）的最小路径和
// dp[j] = grid[i][j] + min(dp[j], dp[j - 1])
func minPathSum1(grid [][]int) int {

	dp := make([]int, len(grid[0]))

	// 开始dp，逐个计算每个坐标的最小路径和
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if j == 0 { // 第一列
				dp[j] = grid[i][j] + dp[j]
			} else if i == 0 { // 第一行
				dp[j] = grid[i][j] + dp[j-1]
			} else {
				dp[j] = grid[i][j] + min(dp[j], dp[j-1])
			}
		}
	}

	// 返回最后一行最后一列的坐标的最小路径和
	return dp[len(grid[0])-1]
}
