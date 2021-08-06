package main

// 547. 省份数量

// 有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，
// 且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。
//
// 省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。
//
// 给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示
// 第 i 个城市和第 j 个城市直接相连，而 isConnected[i][j] = 0 表示二者不直接相连。
//
// 返回矩阵中 省份 的数量。

func main() {

}

// 类似岛屿数量
// 遍历每一个城市，之后深度访问该城市能到达的所有城市
func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected)) // 记录每个城市是否已被访问过
	res := 0
	for i := 0; i < len(visited); i++ {
		if !visited[i] { // 若当前城市还没有被别的城市连接
			res++
			dfs2(isConnected, i, visited)
		}
	}

	return res
}

// i: 遍历第i个城市的所有连接的城市
func dfs2(grid [][]int, i int, visited []bool) {
	if i < 0 || i >= len(grid) {
		return
	}

	visited[i] = true // 标记当前城市被访问过
	for j := 0; j < len(grid[i]); j++ {
		if grid[i][j] == 1 && !visited[j] {
			dfs2(grid, j, visited)
		}
	}
}
