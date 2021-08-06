package main

// 200. 岛屿数量

func main() {

}

// 回溯法
// 类似 offer 12, 13
// Hot100 463, 695
// 等于机器人有几个可以运动的范围
func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}

	num := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				num++
				dfs1(grid, i, j)
			}
		}
	}

	return num
}

func dfs1(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}

	grid[i][j] = '*'

	dfs1(grid, i-1, j)
	dfs1(grid, i+1, j)
	dfs1(grid, i, j-1)
	dfs1(grid, i, j+1)
}
