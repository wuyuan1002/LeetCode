package main

// 48. 旋转图像

// 给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
//
// 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

// 1. 由外而内，逐层遍历，一层一层旋转
// 2. 先上下对折，再沿对角线对折

// rotate48 .
// 由外而内，逐层遍历，一层一层旋转 -- 对于一个n*n的矩阵，[i, j]位置的数字经过旋转后到了[j, n-1-i]位置
func rotate48(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] = matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}
