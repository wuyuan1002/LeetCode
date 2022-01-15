package main

// 42. 接雨水

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

// func main() {
// 	fmt.Println(trap1([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
// }

// 对比Hot100 84
// 双指针 -- 也可用单调栈来解决
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

// 单调递减栈 -- 对于每个点找其左边和右边第一个大于或等于它的点（这样才能构成凹形）
// 类似Hot100 84, 739
func trap1(height []int) int {
	if height == nil {
		return 0
	}

	res := 0                // 总雨水数
	stack := make([]int, 0) // 单调递减栈

	for i, rightH := range height {
		// 若栈不为空且当前元素大于栈顶元素，依次将栈内大于当前元素的值出栈，计算它能接的雨水数量
		for len(stack) > 0 && rightH > height[stack[len(stack)-1]] {
			// 获取当前栈顶元素，并出栈
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 若当前栈顶元素为栈内第一个元素，直接退出循环然后将当前元素入栈，因为左面没有比它高的元素，构不成凹形无法盛接雨水
			if len(stack) == 0 {
				break
			}

			// 获取它左面第一个比它高的元素下标
			left := stack[len(stack)-1]
			curWidth := i - left - 1                             // 左右侧比它高的元素间距
			curHeight := min(height[left], rightH) - height[top] // 左右侧高的元素较低的那个和当前元素的高度差
			res += curWidth * curHeight
		}
		// 若栈为空或当前元素小于栈顶元素了，将当前元素入栈
		stack = append(stack, i)
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
