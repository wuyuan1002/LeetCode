package main

// 221. 最大正方形

// 在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。

// func main() {

// }

// 动态规划
// 某一位置的最大正方形边长 = 它上、左、左上三个位置中边长的最小值+1
func maximalSquare(matrix [][]byte) int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	min := func(a, b byte) byte {
		if a <= b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	maxSide := 0 // 最大变长
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					// 第一行或第一列
					matrix[i][j] = '1'
				} else {
					matrix[i][j] = min(matrix[i-1][j], min(matrix[i][j-1], matrix[i-1][j-1])) + 1
				}

				// 更新最大变长
				maxSide = max(maxSide, int(matrix[i][j]-'0'))
			}
		}
	}

	return maxSide * maxSide
}
