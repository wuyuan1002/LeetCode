package main

import (
	"fmt"
)

// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。

func main() {
	fmt.Println(largestRectangleArea([]int{2, 3, 4, 5}))
}

// 一直解答错误
// 对比Hot100 42
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
		// 若栈为空或当前元素大于等于栈顶元素，直接入栈 -- 栈内元素的右面比它低的元素还没有确定下来
		if len(stack) == 0 || (len(stack) > 0 && heights[i] >= heights[stack[len(stack)-1]]) {
			stack = append(stack, i)
			continue
		}

		// 若当前元素小于栈顶元素，说明有栈内元素的右边界已经确定，
		// 接下来依次弹出栈内元素，当前元素为它的右边界，栈内它的前一个元素是它的左边界，计算他们所能构成的最大矩形面积
		// 直到栈内元素小于了当前元素或栈内元素全部弹出，再将当前元素入栈，让当前元素位于栈顶，成为最大的元素
		j := len(stack) - 1
		for ; j >= 0 && heights[stack[j]] >= heights[i]; j-- {
			left := 0 // 栈内元素左面第一个比它小的下标
			if j == 0 {
				left = -1
			} else {
				left = stack[j-1]
			}
			maxArea = max(maxArea, heights[stack[j]]*(i-left-1))
		}
		// 将j之后的元素去掉，将当前元素入栈 -- 此时当前元素为栈内最大值
		stack = append(stack[:j+1], i)
	}

	// 计算栈中剩余元素所能构成的最大矩形面积，或者是数组本身单调递增时求面积，如heights == []int{2,3,4,5}
	for index := len(stack) - 1; index >= 0; index-- {
		maxArea = max(maxArea, heights[stack[index]]*(len(heights)-stack[index]))
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
