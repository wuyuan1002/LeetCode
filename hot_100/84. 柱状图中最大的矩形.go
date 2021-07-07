package main

import (
	"fmt"
)

// 84. 柱状图中最大的矩形

// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。

func main() {
	// fmt.Println(largestRectangleArea([]int{2, 3, 4, 5}))
	// fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea([]int{2, 1, 2}))
}

// 一直解答错误
// 对比Hot100 42, 739
// 单调栈 -- 单调递增栈
// 某个柱子能构成的最大面积 == 它的高度 * 左右第一个比它低的柱子的坐标差
func largestRectangleArea(heights []int) int {
	if heights == nil || len(heights) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	// 最大面积
	maxArea := 0

	// 单调递增栈 -- 存元素下标，栈内每个元素的前一个元素都是它左面第一个比它低的元素
	stack := make([]int, 0)

	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 当前柱子与左面第一个比它高的柱子的距离 -- 若栈中当前柱子左面没有元素，说明当前柱子左面的元素都比它高，它左面的宽度就是它的下标值-1，否则左面的宽度等于它与左面第一个比它高的元素的距离
			leftWidth := top
			if len(stack) > 0 {
				leftWidth = top - stack[len(stack)-1] - 1
			}
			// 当前柱子与右面第一个比它高的柱子的距离 -- 右面宽度等于它与右面第一个比它高的元素的距离
			rightWidth := i - top - 1

			curWidth := leftWidth + rightWidth + 1 // 总的宽度
			curHeight := heights[top]              // 总的高度
			maxArea = max(maxArea, curHeight*curWidth)
		}

		stack = append(stack, i)
	}

	// 防止栈始终为递增，如[]int{2, 3, 4, 5}
	if len(stack) > 0 {
		// 获取当前栈内的最大高度
		maxHeight := stack[len(stack)-1]
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			curWidth := maxHeight - top + 1
			curHeight := heights[top]

			maxArea = max(maxArea, curHeight*curWidth)
		}
	}

	return maxArea
}

// 解答
func largestRectangleArea2(heights []int) int {

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	monoStack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			right[monoStack[len(monoStack)-1]] = i
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			left[i] = -1
		} else {
			left[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

// 2.遍历每一个柱子，然后向两边扩散求它的最大面积 -- 超出时间限制
func largestRectangleArea1(heights []int) int {
	if heights == nil || len(heights) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	maxArea := 0 // 最大面积
	for i := 0; i < len(heights); i++ {
		area := heights[i] // 当前柱子可以构成的最大面积
		// 向左扩散
		for l := i - 1; l >= 0; l-- {
			if heights[l] < heights[i] {
				break
			}
			area += heights[i]
		}
		// 向右扩散
		for r := i + 1; r < len(heights); r++ {
			if heights[r] < heights[i] {
				break
			}
			area += heights[i]
		}

		// 更新最大面积
		maxArea = max(area, maxArea)
	}

	return maxArea
}
