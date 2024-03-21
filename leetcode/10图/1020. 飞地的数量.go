package main

// 1020. 飞地的数量

// 给你一个大小为 m x n 的二进制矩阵 grid ，其中 0 表示一个海洋单元格、1 表示一个陆地单元格。
// 一次 移动 是指从一个陆地单元格走到另一个相邻（上、下、左、右）的陆地单元格或跨过 grid 的边界。
// 返回网格中 无法 在任意次数的移动中离开网格边界的陆地单元格的数量。

// numEnclaves .
// 同 leetcode 695、130
// 本题相当于是求四周都是水（不靠边界）的陆地总面积
// 遍历地图上的每个坐标，求访问到的陆地面积，同时判断所属陆地是否为飞地（四周都是水，不靠边界）
func numEnclaves(grid [][]int) int {
	// 总结果 -- 四周都是水的陆地总面积
	result := 0

	// 遍历地图的每个坐标，统计飞地陆地总面积
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 若当前节点已被访问过则为-1、水为0、只有当前节点是陆地且未被访问过为1
			if grid[i][j] == 1 {
				// 当前坐标所属陆地是否为飞地 -- 默认为true
				isFlyArea := true
				// 计算当前坐标所属陆地的面积并统计当前陆地是否为飞地
				area := defNumEnclaves(grid, i, j, &isFlyArea)
				// 若为飞地则累加到总结果
				if isFlyArea {
					result += area
				}
			}
		}
	}

	return result
}

// defNumEnclaves .
// 返回给定坐标所属陆地的面积以及当前坐标所属陆地是否为飞地（即四周都是水的陆地）
func defNumEnclaves(grid [][]int, i, j int, isFlyArea *bool) int {
	// 若当前坐标超出地图范围或不是陆地或已被访问过
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != 1 {
		return 0
	}

	// 标记当前位置已被访问
	grid[i][j] = -1

	// 若当前坐标不在地图边界（是飞地）、若当前坐标为地图边界（不是飞地）
	// 若isFlyArea已经是false，说明当前坐标所属陆地已经不是飞地了，不应重复赋值防止覆盖 -- 所以前面要加上&&（false和任何数做与运算都为false）
	*isFlyArea = *isFlyArea && !(i == 0 || i == len(grid)-1 || j == 0 || j == len(grid[0])-1)

	// 返回当前坐标所属陆地总面积
	return 1 + defNumEnclaves(grid, i+1, j, isFlyArea) + defNumEnclaves(grid, i-1, j, isFlyArea) + defNumEnclaves(grid, i, j+1, isFlyArea) + defNumEnclaves(grid, i, j-1, isFlyArea)
}
