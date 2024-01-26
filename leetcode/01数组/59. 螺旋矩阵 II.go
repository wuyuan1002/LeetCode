package main

// 59. 螺旋矩阵 II

// 给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。

// generateMatrix .
func generateMatrix(n int) [][]int {

	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	num, count := 1, n*n                       // 当前数字和目标数字大小
	top, right, bottom, left := 0, n-1, n-1, 0 // 上下左右边界

	for num <= count {
		// 最上面一行
		for i := left; i <= right; i++ {
			result[top][i] = num
			num++
		}
		top++

		// 最右面一列
		for i := top; i <= bottom; i++ {
			result[i][right] = num
			num++
		}
		right--

		// 最下面一行
		for i := right; i >= left; i-- {
			result[bottom][i] = num
			num++
		}
		bottom--

		// 最左面一列
		for i := bottom; i >= top; i-- {
			result[i][left] = num
			num++
		}
		left++
	}

	return result
}
