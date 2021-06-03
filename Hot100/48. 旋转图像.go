package main

// 给定一个 n×n 的二维矩阵matrix 表示一个图像。请你将图像顺时针旋转 90 度。
// 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

func main() {

}

// 1. 由外而内，逐层遍历，一层一层旋转 -- 对于一个n*n的矩阵，(i,j)经过旋转后到了(j,n-1-i)
// 2. 先上下对折，再沿对角线对折

// 由外而内，逐层遍历，一层一层旋转
// 结果不对-------------
func rotate1(matrix [][]int) {
	if matrix == nil || len(matrix) == 0 {
		return
	}

	length := len(matrix)
	for i := 0; i < length/2; i++ { // 由外而内，逐层遍历 -- 当前是第几圈
		for j := i; j < length-i-1; j++ { // 一圈的长度 -- 一圈数字的横坐标
			matrix[i][j], matrix[j][length-i-1], matrix[length-i-1][j], matrix[j][i] =
				matrix[j][i], matrix[i][j], matrix[j][length-i-1], matrix[length-i-1][j]
		}
	}
}

// 结果对-------------
func rotate(matrix [][]int) {

	if matrix == nil || len(matrix) == 0 {
		return
	}

	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}
