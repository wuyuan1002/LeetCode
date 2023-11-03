package main

// 63. 不同路径 II

// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
// 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

// uniquePathsWithObstacles .
// 二维dp
// dp[i][j]表示在[i, j]位置的走法个数
// 到达某一个位置的走法 == 它上面位置的走法个数 + 左面位置的走法个数
// 因为障碍物的阻挡，所以在障碍物处的走法个数为0，表示计算障碍物下方和右侧的走法时，不能从障碍物处过去
// dp[i][j] = dp[i][j-1] + dp[i-1][j]
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid == nil || len(obstacleGrid) == 0 {
		return 0
	}

	// 构造dp数组
	dp := make([][]int, len(obstacleGrid))
	for i := 0; i < len(obstacleGrid); i++ {
		dp[i] = make([]int, len(obstacleGrid[0]))
	}

	// 初始化dp数组 -- 初始化第一行和第一列的走法个数，第一行和第一列在障碍物后面的走法个数都为0
	for i := 0; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 1 {
			// 遇到障碍物直接break，保持障碍物后面的位置走法个数为0
			break
		}
		dp[i][0] = 1
	}
	for j := 0; j < len(obstacleGrid[0]); j++ {
		if obstacleGrid[0][j] == 1 {
			// 遇到障碍物直接break，保持障碍物后面的位置走法个数为0
			break
		}
		dp[0][j] = 1
	}

	// 挨个行遍历数组，计算每一个位置的走法个数
	for i := 1; i < len(obstacleGrid); i++ { // 从第二行开始遍历，因为第一行已经初始化过了
		for j := 1; j < len(obstacleGrid[0]); j++ { // 从第二列开始遍历，因为第一列已经初始化过了
			if obstacleGrid[i][j] == 1 {
				// 遇到障碍物，将障碍物位置的走法个数置为0，表示不能从障碍物处走
				dp[i][j] = 0
			} else {
				// 其它位置的走法个数为上面和左面的走法个数相加
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	// 返回最后一行最后一个位置的走法个数
	return dp[len(dp)-1][len(dp[0])-1]
}

// uniquePathsWithObstacles1 .
// 滚动数组 -- 一维dp
// 从上面的二维dp我们可以发现，每次遍历下一行时，其实只使用前一行的数据，对于前两行前三行等的数据其实已经没用了
// 我们可以使用一维数组滚动更新每一行的走法个数，计算下一行时直接使用当前数组数字计算后赋值即可
// dp[j]表示在某一行第j个位置的走法个数
// dp[j] = dp[j] + dp[j - 1]
func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
	// 构造dp数组 -- dp数组记录的是前一行每个位置的走法个数
	dp := make([]int, len(obstacleGrid[0]))

	// 初始化dp数组 -- 第一行第一个位置的走法个数为1，因为要从[0, 0]位置出发，所以第一个位置一定不会是障碍物
	dp[0] = 1

	// 遍历计算每个位置的走法个数并保存中间结果到dp数组
	for i := 0; i < len(obstacleGrid); i++ { // 遍历每一行
		for j := 0; j < len(obstacleGrid[0]); j++ { // 遍历当前行的每一列
			if obstacleGrid[i][j] == 1 {
				// 当前位置为障碍物，将当前行当前位置的走法个数置为0，表示不能从当前位置走
				dp[j] = 0
			} else if j == 0 {
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
	return dp[len(obstacleGrid[0])-1]
}
