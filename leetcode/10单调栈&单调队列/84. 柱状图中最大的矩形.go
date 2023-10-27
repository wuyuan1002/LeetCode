package main

// 84. 柱状图中最大的矩形

// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。

// largestRectangleArea .
// 单调栈 -- 单调递减栈
// 每个柱子能构成的最大面积 == 当前柱子高度 * 左右比它低的柱子间距离
func largestRectangleArea(heights []int) int {
	// 最大面积
	result := 0
	// 单调递减栈 -- 从栈顶到栈尾单调递减，栈内每个元素的前一个元素都是它左面第一个比它低的元素的下标
	// 栈内元素都是还没有找到它右面比它小的元素的下标，当找到其右面比其小的元素时，开始出栈计算它能够构成的面积
	stack := make([]int, 0)

	// 向数组的最后加入一个比数组内所有元素都小的元素，保证数组内所有元素都一定会有右侧比自己小的值
	// 防止当数组始终为递增时不会触发计算面积（因为触发计算每一列能够构成的面积的时机是找到它的右侧比它小的值的时候），如[]int{2, 3, 4, 5}
	heights = append(heights, -1)

	for i := 0; i < len(heights); i++ {
		// 若当前元素比栈顶元素小，说明找到了栈顶元素的右侧第一个比它小的元素，
		// 开始逐个计算栈内比当前元素大的元素能构成的最大面积，同时更新总结果
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			// 获取当前栈顶元素并出栈
			numIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 获取当前栈顶元素左侧第一个比其小的元素下标
			// 若栈不为空，则栈内下一个元素为当前栈顶元素的左侧第一个比它低的元素下标
			// 若栈为空，则说明当前栈顶元素左面的元素都比它高，则将左侧元素下标置为-1
			leftIndex := -1
			if len(stack) > 0 {
				leftIndex = stack[len(stack)-1]
			}

			// 计算当前栈顶元素能构成的最大面积 -- 当前柱子高度 * 左右比它低的柱子间距离
			numArea := heights[numIndex] * (i - leftIndex - 1)

			// 更新最大矩形面积
			result = max(result, numArea)
		}

		// 若当前元素比栈顶元素大，说明当前元素比栈内元素都大，将当前元素入栈，等待它右侧比它小的元素出现
		stack = append(stack, i)
	}

	return result
}
