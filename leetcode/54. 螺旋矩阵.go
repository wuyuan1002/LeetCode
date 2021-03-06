package main

// 54. 螺旋矩阵

// func main() {

// }

func spiralOrder(matrix [][]int) []int {
	if matrix == nil {
		return nil
	}

	result := make([]int, 0, len(matrix)*len(matrix[0]))
	orderNum := min(len(matrix)+1, len(matrix[0])+1) / 2
	for i := 0; i < orderNum; i++ {
		printOrder(matrix, i, &result)
	}
	return result
}

// 打印一圈
func printOrder(matrix [][]int, start int, result *[]int) {
	endX := len(matrix[0]) - start - 1
	endY := len(matrix) - start - 1

	for i := start; i <= endX; i++ {
		*result = append(*result, matrix[start][i])
	}
	for j := start + 1; j <= endY; j++ {
		*result = append(*result, matrix[j][endX])
	}
	if endY > start {
		for i := endX - 1; i >= start; i-- {
			*result = append(*result, matrix[endY][i])
		}
	}
	if endX > start {
		for j := endY - 1; j > start; j-- {
			*result = append(*result, matrix[j][start])
		}
	}
}

// 第二次做
func spiralOrder1(matrix [][]int) []int {
	if matrix == nil {
		return nil
	}

	// 总结果
	result := make([]int, 0, len(matrix)*len(matrix[0]))
	// 要打印的圈数
	orderNum := min(len(matrix)+1, len(matrix[0])+1) / 2

	for i := 0; i < orderNum; i++ {
		printOrder1(matrix, i, &result)
	}
	return result
}

// 打印一圈
func printOrder1(matrix [][]int, start int, result *[]int) {
	endX := len(matrix[0]) - start - 1
	endY := len(matrix) - start - 1

	for i := start; i <= endX; i++ {
		*result = append(*result, matrix[start][i])
	}
	for j := start + 1; j <= endY; j++ {
		*result = append(*result, matrix[j][endX])
	}
	if endY > start { // 是否只有一行
		for i := endX - 1; i >= start; i-- {
			*result = append(*result, matrix[endY][i])
		}
	}
	if endX > start { // 是否只有一列
		for j := endY - 1; j > start; j-- {
			*result = append(*result, matrix[j][start])
		}
	}
}
