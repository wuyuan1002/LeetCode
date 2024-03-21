package main

// 827. 最大人工岛

// 给你一个大小为 n x n 二进制矩阵 grid 。最多 只能将一格 0 变成 1 。
// 返回执行此操作后，grid 中最大的岛屿面积是多少？
// 岛屿 由一组上、下、左、右四个方向相连的 1 形成。

// largestIsland .
// 1. 遍历地图，统计每个岛屿的面积，并为每个岛屿记录编号，将每个岛屿的面积记录起来
// 2. 遍历每一个海水坐标，将它上下左右连接的岛屿面积相加就是把该点变成陆地之后的人工岛面积
func largestIsland(grid [][]int) int {
	// 使用一个map记录每个岛屿的面积 -- k: 岛屿编号，v: 岛屿面积
	// 因为在地图中，0表示海水、1表示陆地，所以我们将岛屿的编号从2开始，
	// 在深度遍历一块岛屿时，同时将岛屿的编号赋值给岛屿的每一个坐标，这样给定一个坐标就可以得知该坐标属于哪个岛屿了
	landAreaMap := make(map[int]int)

	// 遍历地图的每一个坐标，在地图上标记每个岛屿编号并统计岛屿面积记录到landAreaMap
	landNum := 2 // 岛屿编号 -- 因为0为海水、1为陆地，所以将岛屿的编号从2开始
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				// 若当前坐标是为访问过的陆地，则计算当前岛屿面积并记录到map中，同时递增岛屿编号
				area := dfsLandArea(grid, i, j, landNum)
				landAreaMap[landNum] = area
				landNum++
			}
		}
	}

	// 若地图中只有一个陆地且陆地面积等于整个地图面积，说明整个地图全是陆地，返回整个地图的面积
	// 如 [[1, 1], [1, 0]] 整块地图都是陆地，应返回4，若进行下面的循环的话会因为找不到为0的坐标而导致result为0
	if len(landAreaMap) == 1 && landAreaMap[landNum-1] == len(grid)*len(grid[0]) {
		return len(grid) * len(grid[0])
	}

	// 遍历地图中的每一个海水，统计将它变成陆地后上下左右相连接的岛屿面积，更新最大的人工岛面积
	result := 0 // 总结果 -- 最大的人工岛面积
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				// 用来记录当前坐标已经计算入人工岛面积的岛屿编号，防止重复计算，比如 [[1, 1], [1, 0]]，由于上面和左面都是相同岛屿，会导致岛屿面积被加入两次得到错误结果7，而不是正确结果4
				visited := make(map[int]bool)

				// 当前坐标变成1后的人工岛面积 -- 因为当前坐标也提供1个面积，所以初始值为1
				area := 1

				// 相加当前坐标上下左右相连接的岛屿面积，并记录已连接的岛屿到visited
				if i < len(grid)-1 && grid[i+1][j] != 0 && !visited[grid[i+1][j]] {
					area += landAreaMap[grid[i+1][j]]
					visited[grid[i+1][j]] = true
				}
				if i > 0 && grid[i-1][j] != 0 && !visited[grid[i-1][j]] {
					area += landAreaMap[grid[i-1][j]]
					visited[grid[i-1][j]] = true
				}
				if j < len(grid[0])-1 && grid[i][j+1] != 0 && !visited[grid[i][j+1]] {
					area += landAreaMap[grid[i][j+1]]
					visited[grid[i][j+1]] = true
				}
				if j > 0 && grid[i][j-1] != 0 && !visited[grid[i][j-1]] {
					area += landAreaMap[grid[i][j-1]]
					visited[grid[i][j-1]] = true
				}

				// 更新最大的人工岛面积
				result = max(area, result)
			}
		}
	}

	return result
}

// dfsLandArea .
// 给定一个坐标，深度遍历该坐标所属的岛屿面积，并将该岛屿的各个坐标赋值为当前岛屿的编号，返回该岛屿的面积
// grid：地图、i、j：当前坐标、landNum：新的岛屿编号
func dfsLandArea(grid [][]int, i, j int, landNum int) int {
	// 若超出地图边界或坐标不是未被访问过的陆地则直接返回0
	// grid[i][j] -- 0海水、1未被访问过的陆地、其它大于1的数为已经被访问过的陆地，其值为岛屿的编号
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != 1 {
		return 0
	}

	// 标记当前坐标所属的岛屿编号
	grid[i][j] = landNum

	// 上下左右遍历坐标，返回最终的岛屿面积
	return 1 + dfsLandArea(grid, i+1, j, landNum) + dfsLandArea(grid, i-1, j, landNum) + dfsLandArea(grid, i, j+1, landNum) + dfsLandArea(grid, i, j-1, landNum)
}
