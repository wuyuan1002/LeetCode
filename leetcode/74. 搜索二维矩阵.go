package main

// 74. 搜索二维矩阵

// 编写一个高效的算法来判断m x n矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
// 每行中的整数从左到右按升序排列。
// 每行的第一个整数大于前一行的最后一个整数。

func main() {

}

// 同 240，offer 04
func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 {
		return false
	}

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
