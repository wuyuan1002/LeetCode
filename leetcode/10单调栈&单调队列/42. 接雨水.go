package main

// 42. 接雨水

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

// trap .
// 1. 双指针 -- 类似 11
// 一列一列求，每一列所能接的雨水数量等于当前列高度和它左右最高列相对低的那列的差决定
//
// 2. 单调栈 -- 类似 84、239、739、496、503
// 单调递减栈 -- 对于每个点找其左边和右边第一个大于或等于它的点（这样才能构成凹形）
func trap(height []int) int {

	l, r := 0, len(height)-1 // 定义双指针，始终向中心移动，表示当前正在求的列
	lmax, rmax := 0, 0       // 当前列的左右最大高度
	result := 0              // 总结果

	for l < r {
		// 向前移动了l或r，重新计算当前所求列的左右最大高度
		// 若新指向下标的值都比左右最大值小，说明所指向的列会盛水，计算后继续移动l或r
		// 若新指向下标的值本身更新了左右的最大值，说明所指向的那一列不会盛水，在后面求的时候差值会是0
		lmax = max(lmax, height[l])
		rmax = max(rmax, height[r])

		if lmax > rmax {
			// 若左侧最大高度比右侧最大高度高，说明当前列的盛水量为当前列与右侧最大高度的差值（每一列的宽度为1，所以不用乘宽度了）
			// 说明当前正在求的列为r指针所指向的列，此时l指针指向的为lmax没变
			result += rmax - height[r]
			// 向前移动r指针，r指针所指向的为下一个要求的列 -- 每次都是移动数字小的那个
			r--
		} else {
			// 若右侧最大高度比左侧最大高度高，说明当前列的盛水量为当前列与左侧最大高度的差值（每一列的宽度为1，所以不用乘宽度了）
			// 向后移动l指针，l指针所指向的为下一个要求的列
			result += lmax - height[l]
			// 向后移动l指针，l指针所指向的为下一个要求的列 -- 每次都是移动数字小的那个
			l++
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
