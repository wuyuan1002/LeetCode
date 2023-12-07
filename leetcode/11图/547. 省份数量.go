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

// findCircleNum .
// 同 leetcode 200、797
// 遍历每一个城市，之后深度访问该城市能到达的所有城市，同时使用visited记录城市是否已被访问，并累加省份数量
// isConnected是一个 n x n 的矩阵，len(isConnected)表示城市数量
func findCircleNum(isConnected [][]int) int {
	// 总结果 -- 省份数量
	result := 0

	// 使用visited数组用来保存各个城市是否已被访问过
	visited := make([]bool, len(isConnected))

	// 遍历每一个城市，将所有它连接的城市进行标记并累加省份数量
	for i := 0; i < len(isConnected); i++ {
		if !visited[i] {
			// 若当前城市还未被其它城市连接，则说明该城市处于一个新的省份，深度标记当前城市所连接的其它城市并将省份数量+1
			result++
			dfsFindCircleNum(isConnected, i, visited)
		}
	}

	return result
}

// dfsFindCircleNum 深度访问第i个城市能到达的所有城市，并把已访问的城市标记为已访问
func dfsFindCircleNum(isConnected [][]int, i int, visited []bool) {
	if i < 0 || i > len(isConnected)-1 {
		return
	}

	// 标记当前城市已被访问
	visited[i] = true

	// 遍历所有城市，寻找与当前城市相连接的城市并进行深度遍历
	for j := 0; j < len(isConnected); j++ {
		if isConnected[i][j] == 1 && !visited[j] {
			dfsFindCircleNum(isConnected, j, visited)
		}
	}
}
