package main

// 463. 岛屿的周长

// 给定一个 row x col 的二维网格地图 grid ，其中：grid[i][j] = 1 表示陆地，
// grid[i][j] = 0 表示水域。
//
// 网格中的格子 水平和垂直 方向相连（对角线方向不相连）。整个网格被水完全包围，
// 但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。
//
// 岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。
// 网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿的周长。

func main() {

}

// 回溯法
// 类似 offer 12, 13
func islandPerimeter(grid [][]int) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				// 题目说只会有一个岛屿，因此找到一个点是岛屿时，直接遍历并返回即可
				return dfs10(grid, i, j)
			}
		}
	}

	return 0
}

// 从土地到土地，不会跨越边长，土地到海洋，土地到超出边界，都会跨越一条边长
func dfs10(grid [][]int, i, j int) int {
	if i < 0 || i > len(grid)-1 || j < 0 || j > len(grid[0])-1 || grid[i][j] == 0 {
		// 若当前位置是海洋或超出边界，土地到海洋，土地到超出边界，返回1
		return 1
	} else if grid[i][j] == -1 {
		// 当前位置已被访问过，土地到土地，不跨跃边长，返回0
		return 0
	}

	// 标记当前位置已被访问
	grid[i][j] = -1

	// 遍历前后左右
	return dfs10(grid, i+1, j) +
		dfs10(grid, i-1, j) +
		dfs10(grid, i, j+1) +
		dfs10(grid, i, j-1)
}
