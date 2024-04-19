package main

// 42. 接雨水

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

// trap1 .
// 1. 双指针（按列求） -- 同 leetcode 11. 盛最多水的容器
// 一列一列求，每一列所能接的雨水数量等于当前列高度和它左右最高列相对低的那列的差决定（每一列的宽度为1，所以不用乘宽度了）
func trap1(height []int) int {
	l, r := 0, len(height)-1 // 双指针，始终向中心移动，表示当前正在求的列
	lmax, rmax := 0, 0       // 当前列的左右最大高度
	result := 0              // 总结果

	for l < r {
		// 更新左右侧的最大高度
		lmax = max(lmax, height[l])
		rmax = max(rmax, height[r])

		// 计算当前列的盛水量并加到总结果中 -- 每次都向前移动小的那个指针，防止漏掉更新小的那一侧的最大值
		if lmax > rmax {
			// 若右侧比左侧小，则遍历右指针，因为不用担心左面会漏
			result += rmax - height[r]
			r--
		} else {
			// 若左侧比右侧小，则遍历左指针，因为不用担心右面会漏
			result += lmax - height[l]
			l++
		}
	}

	return result
}

// trap2 .
// 2. 单调栈（按行求） -- 类似 84、239、739、496、503
// 单调递减栈 -- 对于每个点找其左边和右边第一个大于或等于它的点（这样才能构成凹形）
func trap2(height []int) int {
	// 总结果
	result := 0
	// 单调递减栈 -- 栈内元素从栈底到栈顶单调递减，栈顶元素为栈内最小值，用于存还没有出现凹槽的列的下标，
	// 也就是栈内元素只找到了左侧比它大的值，还没有找到右侧比它大的值，还没有形成凹槽
	stack := make([]int, 0)

	for i := 0; i < len(height); i++ {
		// 若栈不为空且当前元素比栈顶元素大，说明栈内有元素已经找到了它们右侧比自己高的值，形成了凹槽
		// 开始计算栈内比当前值小的元素的盛水量
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			// 获取当前栈顶元素并出栈
			numIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 若当前栈顶元素为栈内第一个元素，直接退出循环然后将当前元素入栈，因为左面没有比它高的元素，构不成凹形无法盛接雨水
			if len(stack) == 0 {
				break
			}

			// 计算当前栈顶元素的盛水量
			left := height[stack[len(stack)-1]]           // 左侧高列高度
			right := height[i]                            // 右侧高列高度
			width := i - stack[len(stack)-1] - 1          // 左右侧高列间距
			height := min(left, right) - height[numIndex] // 较低的左右侧高列与当前栈顶元素的高度差

			// 计算当前栈顶元素的盛水量并加到总结果
			result += width * height
		}

		// 若栈为空或当前元素小于栈顶元素了，将当前元素入栈，等待当前元素的右侧高列出现
		stack = append(stack, i)
	}

	// 遍历完所有列后，栈内元素都是只有左侧高列，没有右侧高列的元素，
	// 即它们没有形成凹槽，没办法盛接雨水
	return result
}

// max .
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min .
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
