package main

// 54. 螺旋矩阵

// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

// spiralOrder .
func spiralOrder(matrix [][]int) []int {
	result := make([]int, 0, len(matrix)*len(matrix[0])) // 总结果
	orderNum := min(len(matrix)+1, len(matrix[0])+1) / 2 // 要打印的圈数

	// 打印每一圈
	for i := 0; i < orderNum; i++ {
		printOrder(matrix, i, &result)
	}

	return result
}

// printOrder 打印一圈
func printOrder(matrix [][]int, start int, result *[]int) {
	endX := len(matrix[0]) - start - 1 // 该圈最大横坐标
	endY := len(matrix) - start - 1    // 该圈最大纵坐标

	// 打印上面行
	for j := start; j <= endX; j++ {
		*result = append(*result, matrix[start][j])
	}

	// 打印右面列
	for i := start + 1; i <= endY; i++ {
		*result = append(*result, matrix[i][endX])
	}

	// 打印下面行 -- 需要判断下是否只有一行，防止只有一行时被打印多次
	if endY > start {
		for j := endX - 1; j >= start; j-- {
			*result = append(*result, matrix[endY][j])
		}
	}

	// 打印左面列 -- 需要判断下是否只有一列，防止只有一列时被打印多次
	if endX > start {
		for i := endY - 1; i > start; i-- {
			*result = append(*result, matrix[i][start])
		}
	}
}
