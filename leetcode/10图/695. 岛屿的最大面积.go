package main

// 695. 岛屿的最大面积

// 给定一个包含了一些 0 和 1 的非空二维数组grid 。
//
// 一个岛屿是由一些相邻的1(代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。
// 你可以假设grid 的四个边缘都被 0（代表水）包围着。
//
// 找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)

// maxAreaOfIsland .
func maxAreaOfIsland(grid [][]int) int {
	// 总结果 -- 最大陆地面积
	result := 0

	// 挨个遍历二维数组的每个位置，判断是否为水、陆地、或已被访问过，
	// 若为未访问过的陆地，则访问该陆地计算该陆地的面积，并更新最大的陆地面积
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 若当前节点已被访问过则为-1、水为0、只有当前节点是陆地且未被访问过为1
			if grid[i][j] == 1 {
				// 计算当前节点所属的面积并更新最大的陆地面积
				area := defMaxAreaOfIsland(grid, i, j)
				result = max(result, area)
			}
		}
	}

	// 返回总结果
	return result
}

// defMaxAreaOfIsland 给定一个坐标，上线左右进行遍历，计算并返回该坐标所属陆地的面积
func defMaxAreaOfIsland(grid [][]int, i, j int) int {
	// 若当前位置不是岛屿或已被访问过，直接返回
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != 1 {
		return 0
	}

	// 标记当前位置已被访问
	grid[i][j] = -1

	// 返回当前坐标所属陆地的面积 -- 当前坐标占地为1 + 上下左右坐标的占地面积
	return 1 + defMaxAreaOfIsland(grid, i+1, j) + defMaxAreaOfIsland(grid, i-1, j) + defMaxAreaOfIsland(grid, i, j+1) + defMaxAreaOfIsland(grid, i, j-1)
}

// max .
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
