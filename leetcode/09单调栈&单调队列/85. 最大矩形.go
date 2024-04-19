package main

// 85. 最大矩形

// 给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。

// maximalRectangle .
// 同 leetcode 84. 柱状图中最大的矩形
//
// 分别将matrix的每一行作为横坐标基线，从横坐标基线向上连续1的个数作为该位置的值（即高度），此时便将问题转换为84题的解法，求出以该横坐标为基线所能构成的最大举行面积，
// 然后分别计算以matrix的每一个横坐标作为基线能够得到的最大举行面积，在这些面积中取最大值便是整个matrix所能构成的最大矩形面积
//
// 此题其实就是：求二维数组的最大矩形面积 -> 求多个坐标系中的最大矩形面积 -> 求一个坐标系中的最大矩形面积 -> 多个坐标系的最大面积取最大值就是二维数组的最大矩形面积
//
// 单调栈 -- 单调递增栈
// 每个柱子能构成的最大面积 == 当前柱子高度 * 左右比它低的柱子间距离
func maximalRectangle(matrix [][]byte) int {
	// 用来存以matrix的每一行作为基线横坐标的坐标系，以及其每一列的向上连续1的个数
	allHeights := make([][]int, len(matrix))

	// 分别统计以matrix的每一行作为基线横坐标，其每一列的向上连续1的个数
	for i := 0; i < len(matrix); i++ {
		allHeights[i] = make([]int, len(matrix[0]))

		for j := 0; j < len(matrix[0]); j++ {
			if i == 0 { // 第一行
				allHeights[i][j] = int(matrix[i][j] - '0')
			} else { // 其它行
				if matrix[i][j] == '0' {
					allHeights[i][j] = 0
				} else {
					allHeights[i][j] = allHeights[i-1][j] + 1
				}
			}
		}
	}

	// 分别计算每个坐标系的最大矩形面积 -- 此处与 leetcode 84. 柱状图中最大的矩形 完全一致
	result := 0
	for _, heights := range allHeights {
		result = max(result, largestRectangleArea(heights))
	}

	// 返回所有坐标系中最大的矩形面积
	return result
}
