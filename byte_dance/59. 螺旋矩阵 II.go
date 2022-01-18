package main

// 59. 螺旋矩阵 II

// 给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，
// 且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。

// func main() {
// 	fmt.Println(generateMatrix(3))
// }

func generateMatrix(n int) [][]int {
	if n <= 0 {
		return nil
	}

	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	num, count := 1, n*n         // 当前数字和目标数字大小
	t, r, b, l := 0, n-1, n-1, 0 // 上下左右边界

	for num <= count {
		// 最上面一行
		for i := l; i <= r; i++ {
			result[t][i] = num
			num++
		}
		t++

		// 最右面一列
		for i := t; i <= b; i++ {
			result[i][r] = num
			num++
		}
		r--

		// 最下面一行
		for i := r; i >= l; i-- {
			result[b][i] = num
			num++
		}
		b--

		// 最左面一列
		for i := b; i >= t; i-- {
			result[i][l] = num
			num++
		}
		l++
	}

	return result
}

// 第二次做
func generateMatrix1(n int) [][]int {
	if n <= 0 {
		return nil
	}

	matrix := make([][]int, n)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, n)
	}

	num, count := 1, n*n
	t, r, b, l := 0, n-1, n-1, 0
	for num <= count {
		// 打印上面一行
		for i := l; i <= r; i++ {
			matrix[t][i] = num
			num++
		}
		t++

		// 打印右面一列
		for i := t; i <= b; i++ {
			matrix[i][r] = num
			num++
		}
		r--

		// 打印下面一行
		for i := r; i >= l; i-- {
			matrix[b][i] = num
			num++
		}
		b--

		// 打印左面一列
		for i := b; i >= t; i-- {
			matrix[i][l] = num
			num++
		}
		l++
	}
	return matrix
}
