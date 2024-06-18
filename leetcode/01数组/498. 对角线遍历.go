package main

// 498. 对角线遍历

// 给你一个大小为 m x n 的矩阵 mat ，请以对角线遍历的顺序，用一个数组返回这个矩阵中的所有元素。

// findDiagonalOrder .
// 以第一行和最后一列的每个数字作为起点，不断向左下角遍历构建该条对角线的中间结果
// 然后构建答案时正反方向交替把每个中间结果中的元素加入到答案中
//
// 以第一行和最后一列的每个数字作为起点，遍历偶数时的对角线结果，从右上往左下遍历该对角线，遍历每个数字的左下角数字
// 以第一列和最后一行的每个数字作为起点，遍历奇数时的对角线结果，从左下往右上遍历该对角线，遍历每个数字的右上角数字
func findDiagonalOrder(mat [][]int) []int {
	// 总结果
	result := make([]int, 0, len(mat)*len(mat[0]))

	// i表示当前是遍历到第几条对角线了
	for i := 0; i < len(mat)+len(mat[0])-1; i++ {
		if i%2 == 1 {
			x := max(i-len(mat[0])+1, 0)
			y := min(i, len(mat[0])-1)
			for x < len(mat) && y >= 0 {
				result = append(result, mat[x][y])
				x++
				y--
			}
		} else {
			x := min(i, len(mat)-1)
			y := max(i-len(mat)+1, 0)
			for x >= 0 && y < len(mat[0]) {
				result = append(result, mat[x][y])
				x--
				y++
			}
		}
	}
	return result
}
