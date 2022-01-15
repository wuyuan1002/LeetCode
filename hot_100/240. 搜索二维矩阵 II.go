package main

// 240. 搜索二维矩阵 II

// 编写一个高效的算法来搜索mxn矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
//
// 每行的元素从左到右升序排列。
// 每列的元素从上到下升序排列。

// func main() {

// }

// 同offer 04
func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || matrix[0] == nil || len(matrix[0]) == 0 {
		return false
	}

	i, j := len(matrix)-1, 0 // i,j指向左下角元素，由左下开始移动
	for i >= 0 && j < len(matrix[0]) {
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
