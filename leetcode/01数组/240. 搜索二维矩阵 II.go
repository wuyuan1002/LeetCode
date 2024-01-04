package main

// 240. 搜索二维矩阵 II

// 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
// 每行的元素从左到右升序排列。
// 每列的元素从上到下升序排列。

// searchMatrix240 .
// 从左下角往右上角进行寻找
func searchMatrix240(matrix [][]int, target int) bool {
	i, j := len(matrix)-1, 0
	for i >= 0 && j <= len(matrix[0])-1 {
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			return true
		}
	}

	return false
}
