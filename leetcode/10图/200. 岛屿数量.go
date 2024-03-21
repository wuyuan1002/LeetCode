package main

// 200. 岛屿数量

// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
// 此外，你可以假设该网格的四条边均被水包围。

// numIslands .
func numIslands(grid [][]byte) int {
	// 总结果 -- 陆地总数
	result := 0

	// 挨个遍历二维数组的每个位置，判断是否为水、陆地、或已被访问过，
	// 若为未访问过的陆地，则访问该陆地并标记为已被访问，同时记录岛屿总数+1
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 若当前节点已被访问过则为*、水为0、只有当前节点是陆地且未被访问过为1
			if grid[i][j] == '1' {
				result++
				dfsNumIslands(grid, i, j)
			}
		}
	}

	// 返回总结果
	return result
}

// dfsNumIslands 给定一个坐标，上下左右遍历寻找该点所属的岛屿边界，并标记该岛屿为已访问为*
func dfsNumIslands(grid [][]byte, i, j int) {
	// 若当前位置不是岛屿或已被访问过，直接返回
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}

	// 标记当前位置已被访问
	grid[i][j] = '*'

	// 遍历前后左右标记当前岛屿的边界
	dfsNumIslands(grid, i+1, j)
	dfsNumIslands(grid, i-1, j)
	dfsNumIslands(grid, i, j+1)
	dfsNumIslands(grid, i, j-1)
}
