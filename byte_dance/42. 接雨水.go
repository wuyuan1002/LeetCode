package main

import (
	"fmt"
)

// 42. 接雨水

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

func main() {
	fmt.Println(trap2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

// 1. 双指针
// 2. 单调栈

// 1. 双指针
// 一列一列求，每一列所能接的雨水数量等于当前列高度和它左右最高列相对低的那列的差决定
func trap1(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	l, r := 0, len(height)-1 // 左右指针
	lmax, rmax := 0, 0       // 左右侧的最大值
	res := 0                 // 总雨水量

	for l < r {
		lmax = max(lmax, height[l])
		rmax = max(rmax, height[r])

		if lmax > rmax {
			res += rmax - height[r]
			r--
		} else {
			res += lmax - height[l]
			l++
		}
	}

	return res
}

// 2. 单调栈
// 单调递减栈 -- 对于每个点找其左边和右边第一个大于或等于它的点（这样才能构成凹形）
func trap2(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}

	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}

	stack := make([]int, 0) // 单调递减栈 -- 存下标
	res := 0                // 总雨水量

	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			// 获取并弹出栈顶元素
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				break
			}

			left := stack[len(stack)-1]                 // 左侧第一个高于当前的柱子
			width := i - left - 1                       // 宽度
			heigh := min(height[left], h) - height[top] // 高度

			res += width * heigh // 计算当前柱子的雨水量
		}

		stack = append(stack, i)
	}

	return res
}
