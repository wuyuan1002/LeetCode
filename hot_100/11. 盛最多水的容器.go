package main

// 11. 盛最多水的容器

// 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点(i,ai) 。
// 在坐标内画 n 条垂直线，垂直线 i的两个端点分别为(i,ai) 和 (i, 0) 。
// 找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
//
// 说明：你不能倾斜容器。

// func main() {

// }

// 双指针法，两个指针向中间移动，每次都是移动数字小的那个
func maxArea(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	l, r := 0, len(height)-1 // 左右指针
	res := 0                 // 最大盛水量

	for l < r {
		if height[l] < height[r] {
			res = max(res, height[l]*(r-l))
			l++
		} else {
			res = max(res, height[r]*(r-l))
			r--
		}
	}
	return res
}
