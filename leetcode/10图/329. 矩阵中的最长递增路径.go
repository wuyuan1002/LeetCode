package main

// 329. 矩阵中的最长递增路径

// 给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。
//
// 对于每个单元格，你可以往上，下，左，右四个方向移动。 你 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。

// longestIncreasingPath .
// 深度优先搜索
// 计算从图中每一个节点出发所能得到的最长递增路径长度，这些路径中的最大值即为最终答案，
// 过程中可以使用一个变量store对已经计算过的节点进行缓存，防止对同一个节点进行多次计算
func longestIncreasingPath(matrix [][]int) int {
	// 构建一个二维数组用于缓存从对应节点出发的最长递增路径的长度 -- 初始值都为0表示该节点还未进行计算
	store := make([][]int, len(matrix))
	for i := range store {
		store[i] = make([]int, len(matrix[0]))
	}

	result := 0
	for i := range matrix {
		for j := range matrix[0] {
			result = max(result, dfsLongestIncreasingPath(matrix, store, i, j))
		}
	}

	return result
}

// dfsLongestIncreasingPath .
// 计算从 matrix[i][j] 出发的最长递增路径的长度
func dfsLongestIncreasingPath(matrix [][]int, store [][]int, i, j int) int {
	// 若当前节点已计算过则直接返回
	if store[i][j] != 0 {
		return store[i][j]
	}

	// 递增当前节点出发的最长递增路径的长度为1 -- 至少为1因为当前节点本身贡献一个长度
	store[i][j]++

	// 计算上下左右节点的最长路径长度
	for _, dir := range dirs {
		row, col := i+dir[0], j+dir[1]

		if row >= 0 && row <= len(matrix)-1 && col >= 0 && col <= len(matrix[0])-1 && matrix[row][col] > matrix[i][j] {
			store[i][j] = max(store[i][j], dfsLongestIncreasingPath(matrix, store, row, col)+1)
		}
	}

	return store[i][j]
}

// 用来方便计算上下左右节点
var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
