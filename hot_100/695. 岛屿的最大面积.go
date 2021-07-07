package main

import (
	"fmt"
)

// 695. 岛屿的最大面积

// 给定一个包含了一些 0 和 1 的非空二维数组grid 。
//
// 一个岛屿是由一些相邻的1(代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。
// 你可以假设grid 的四个边缘都被 0（代表水）包围着。
//
// 找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)

func main() {
	fmt.Println(maxAreaOfIsland([][]int{{1, 1, 0}, {1, 0, 0}, {0, 1, 0}}))
}

// 回溯法
// 类似 offer 12, 13
// Hot100 200, 463
// 等于在所有机器人的运动的范围中最大的那个
func maxAreaOfIsland(grid [][]int) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 若当前节点已被访问过则是-1, 水是0, 只有当前节点是陆地时是1
			if grid[i][j] == 1 {
				area := dfs9(grid, i, j)
				maxArea = max(maxArea, area)
			}
		}
	}

	return maxArea
}

func dfs9(grid [][]int, i, j int) int {
	// 若当前位置不是岛屿或已被访问过，直接返回
	if i < 0 || i > len(grid)-1 || j < 0 || j > len(grid[0])-1 || grid[i][j] != 1 {
		return 0
	}

	// 标记当前位置已被访问
	grid[i][j] = -1

	// 前后左右
	return 1 +
		dfs9(grid, i+1, j) +
		dfs9(grid, i-1, j) +
		dfs9(grid, i, j+1) +
		dfs9(grid, i, j-1)
}
