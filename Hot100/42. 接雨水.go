package main

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

func main() {

}

// 双指针
// 一列一列求，每一列所能接的雨水数量等于当前列高度和它左右最高列相对低的那列的差决定
func trap(height []int) int {
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
	lmax, rmax := 0, 0       // 左侧和右侧最高的列
	res := 0                 // 总雨水数

	for l < r {
		// 更新左右侧最大值
		lmax = max(height[l], lmax)
		rmax = max(height[r], rmax)

		if lmax <= rmax {
			// 若左侧比右侧小，则遍历左指针，因为不用担心右面会漏
			res += lmax - height[l]
			l++
		} else {
			// 若右侧比左侧小，则遍历右指针，因为不用担心左面会漏
			res += rmax - height[r]
			r--
		}
	}

	return res
}
