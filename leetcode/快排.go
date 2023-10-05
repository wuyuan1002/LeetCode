package leetcode

// 快排
// 如何优化排序算法: https://mp.weixin.qq.com/s/5HqfRGqPyAhFt0krPgMHOQ

// quickSortWithRecursion 递归快排
func quickSortWithRecursion(nums []int, left, right int) {
	if nums == nil || len(nums) == 0 || left >= right {
		return
	}

	l, r := left, right
	temp := nums[left]
	for l < r {
		for l < r && nums[r] >= temp {
			r--
		}
		for l < r && nums[l] <= temp {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[left], nums[l] = nums[l], nums[left]

	quickSortWithRecursion(nums, left, l-1)
	quickSortWithRecursion(nums, r+1, right)
}

// quickSortWithoutRecursion 非递归快排
func quickSortWithoutRecursion(nums []int, left, right int) {
	stack := []int{left, right} // 使用栈保存每次需要递归的左右界限
	for len(stack) > 0 {
		// 获取本次partition的左右界限
		right = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		left = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 开始partition
		l, r := left, right
		temp := nums[left]
		for l < r {
			for l < r && nums[r] >= temp {
				r--
			}
			for l < r && nums[l] <= temp {
				l++
			}
			if l < r {
				nums[l], nums[r] = nums[r], nums[l]
			}
		}
		nums[left], nums[l] = nums[l], nums[left]

		// 将下一次需要partition的左右边界入栈
		stack = append(stack, left, l)
		stack = append(stack, r+1, right)
	}
}
