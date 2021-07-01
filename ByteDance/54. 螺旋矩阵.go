package main

import (
	"fmt"
)

// 54. 螺旋矩阵

// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}}))
}

func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 {
		return nil
	}

	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}

	result := make([]int, 0)                           // 总结果
	orders := min(len(matrix)+1, len(matrix[0])+1) / 2 // 总共能打印多少圈
	for i := 0; i < orders; i++ {                      // 打印每一圈
		printCircle(matrix, i, &result)
	}
	return result
}

// 打印一圈
func printCircle(matrix [][]int, start int, result *[]int) {
	endX := len(matrix[0]) - start - 1 // 该圈的最远横坐标
	endY := len(matrix) - start - 1    // 该圈的最远纵坐标

	for i := start; i <= endX; i++ { // 打印上面的一行
		*result = append(*result, matrix[start][i])
	}
	for i := start + 1; i <= endY; i++ { // 打印右面的一列
		*result = append(*result, matrix[i][endX])
	}
	if endY-start > 0 {
		for i := endX - 1; i >= start; i-- { // 打印下面的一行
			*result = append(*result, matrix[endY][i])
		}
	}
	if endX-start > 0 {
		for i := endY - 1; i >= start+1; i-- { // 打印左面的一列
			*result = append(*result, matrix[i][start])
		}
	}
}
