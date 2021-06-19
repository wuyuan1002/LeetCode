package main

// 200. 岛屿数量

// 给你一个由'1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
// 此外，你可以假设该网格的四条边均被水包围。

func main() {

}

// 回溯法
// 类似 offer 12, 13
// Hot100 463, 695
// 等于机器人有几个可以运动的范围
func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 若当前节点已被访问过则是*, 水是0, 只有当前节点是陆地时是1
			if grid[i][j] == '1' {
				count++
				dfs8(grid, i, j)
			}
		}
	}

	return count
}

func dfs8(grid [][]byte, i, j int) {
	// 若当前位置不是岛屿或已被访问过，直接返回
	if i < 0 || i > len(grid)-1 || j < 0 || j > len(grid[0])-1 || grid[i][j] != '1' {
		return
	}

	// 标记当前位置已访问过
	grid[i][j] = '*'

	// 遍历前后左右
	dfs8(grid, i+1, j)
	dfs8(grid, i-1, j)
	dfs8(grid, i, j+1)
	dfs8(grid, i, j-1)
}
