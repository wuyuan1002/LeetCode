package main

// 11. 盛最多水的容器

// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。
//
// 说明：你不能倾斜容器。

// maxArea .
// 同 leetcode 42. 接雨水
// 双指针 -- 两个指针向中间移动，每次都是移动数字小的那个
// 左右两个指针用来选择盛水所用的两边柱子，每两个柱子的盛水量 == 两柱子的距离 * 两柱子的中短的那个的高度，
// 每计算完两个柱子的盛水量后更新总结果的最大值，然后向前移动短的那个柱子（因为不用担心另一边会漏），继续计算下两个柱子的盛水脸
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	result := 0

	for l < r {
		if height[l] < height[r] {
			// 左柱子比右柱子低，当前两个柱子盛水量的高度由左柱子决定，计算盛水量后向前移动左柱子
			result = max(result, height[l]*(r-l))
			l++
		} else {
			// 右柱子比左柱子低，当前两个柱子盛水量的高度由右柱子决定，计算盛水量后向前移动右柱子
			result = max(result, height[r]*(r-l))
			r--
		}
	}

	return result
}

// max .
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
