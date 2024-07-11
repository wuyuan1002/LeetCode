package main

// 74. 搜索二维矩阵

// 给你一个满足下述两条属性的 m x n 整数矩阵：
//
// 每行中的整数从左到右按非严格递增顺序排列。
// 每行的第一个整数大于前一行的最后一个整数。
// 给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。

// searchMatrix .
// 数组从上到下，从左至右都是递增的，所以从左下角往右上角进行寻找
// 比目标数字大则向上移动，比目标数字小则向右移动
func searchMatrix(matrix [][]int, target int) bool {
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
