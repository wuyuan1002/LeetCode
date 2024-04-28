package main

// 剑指 Offer 29. 顺时针打印矩阵

// 输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

// func main() {

// }

func spiralOrder(matrix [][]int) []int {
	if matrix == nil {
		return nil
	}
	rows := 0
	lens := len(matrix)
	if lens != 0 {
		rows = len(matrix[0])
	}
	if lens == 0 || rows == 0 {
		return nil
	}

	show := make([]int, 0)
	start := 0 // 开始坐标
	for ; lens > start*2 && rows > start*2; start++ {
		show = printOneCircle(matrix, rows, lens, start, show)
	}
	return show
}

// printOneCircle 打印一圈
func printOneCircle(matrix [][]int, rows int, lens int, start int, show []int) []int {
	endX := rows - 1 - start
	endY := lens - 1 - start
	for i := start; i <= endX; i++ { // 从左到右打印一行
		show = append(show, matrix[start][i])
	}
	if start < endY { // 从上到下打印一列
		for j := start + 1; j <= endY; j++ {
			show = append(show, matrix[j][endX])
		}
	}
	if start < endX && start < endY { // 从右到左打印一行
		for i := endX - 1; i >= start; i-- {
			show = append(show, matrix[endY][i])
		}
	}
	if start < endX && start < endY-1 { // 从下到上打印一列
		for j := endY - 1; j >= start+1; j-- {
			show = append(show, matrix[j][start])
		}
	}
	return show
}

// 第二次做
func spiralOrder1(matrix [][]int) []int {
	if matrix == nil {
		return nil
	}
	rows := 0
	lens := len(matrix)
	if lens > 0 {
		rows = len(matrix[0])
	}
	if lens == 0 || rows == 0 {
		return nil
	}

	result := make([]int, 0) // 保存所有打印结果
	start := 0               // 一圈的开始横纵坐标 -- 左上角坐标
	for start*2 < lens && start*2 < rows {
		printOneCircle1(matrix, rows, lens, start, &result)
		start++
	}
	return result
}
func printOneCircle1(matrix [][]int, rows int, lens int, start int, result *[]int) {
	endx := rows - 1 - start // 一圈的最大横坐标
	endy := lens - 1 - start // 一圈的最大纵坐标

	for i := start; i <= endx; i++ {
		*result = append(*result, matrix[start][i])
	}

	if start < endy { // 判断是不是只有一行
		for j := start + 1; j <= endy; j++ {
			*result = append(*result, matrix[j][endx])
		}
	}

	if start < endx && start < endy {
		for i := endx - 1; i >= start; i-- {
			*result = append(*result, matrix[endy][i])
		}
	}

	if start < endx && start < endy-1 {
		for j := endy - 1; j > start; j-- {
			*result = append(*result, matrix[j][start])
		}
	}
}
