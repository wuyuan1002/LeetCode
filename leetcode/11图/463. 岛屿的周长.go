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

// islandPerimeter .
func islandPerimeter(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				// 题目说只会有一个岛屿, 因此找到一个点是岛屿时, 直接遍历返回即可
				return dfsIslandPerimeter(grid, i, j)
			}
		}
	}

	return 0
}

// dfsIslandPerimeter 求给定坐标所属岛屿的周长
// 土地到土地 -- 不跨越边长
// 土地到海洋、土地到超出边界 -- 跨越一条边长
func dfsIslandPerimeter(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
		// 当前位置是海洋或超出边界，土地到海洋、土地到超出边界 -- 跨越一条边长
		return 1
	} else if grid[i][j] == -1 {
		// 当前位置已被访问过，土地到土地 -- 不跨越边长
		return 0
	}

	// 标记当前位置已被访问
	grid[i][j] = -1

	// 返回上下左右遍历后整个岛屿的周长
	return dfsIslandPerimeter(grid, i+1, j) + dfsIslandPerimeter(grid, i-1, j) + dfsIslandPerimeter(grid, i, j+1) + dfsIslandPerimeter(grid, i, j-1)
}
