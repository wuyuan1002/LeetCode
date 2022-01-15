package main

// 239. 滑动窗口最大值

// 给你一个整数数组 nums，有一个大小为k的滑动窗口从数组的最左侧移动到数组的最右侧。
// 你只可以看到在滑动窗口内的 k个数字。滑动窗口每次只向右移动一位。
//
// 返回滑动窗口中的最大值。

// func main() {
// 	fmt.Println(maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))
// }

// 同 offer 59
func maxSlidingWindow(nums []int, k int) []int {
	if nums == nil || len(nums) == 0 || k <= 0 || k > len(nums) {
		return nil
	}

	result := make([]int, 0)   // 结果集
	maxqueue := make([]int, 0) // 单调递减队列 -- 存下标

	i := 0
	// 移动i形成第一个窗口
	for ; i < k; i++ {
		for len(maxqueue) > 0 && nums[i] >= nums[maxqueue[len(maxqueue)-1]] {
			maxqueue = maxqueue[:len(maxqueue)-1]
		}
		maxqueue = append(maxqueue, i)
	}

	// 将第一个窗口的最大值入结果集
	result = append(result, nums[maxqueue[0]])

	// 向移动窗口
	for ; i < len(nums); i++ {
		// 1. 若出窗口的是最大值，则先将它移出队列
		if i-maxqueue[0]+1 > k {
			maxqueue = maxqueue[1:]
		}

		// 2. 将新入窗口的数据加入到队列中
		for len(maxqueue) > 0 && nums[i] >= nums[maxqueue[len(maxqueue)-1]] {
			maxqueue = maxqueue[:len(maxqueue)-1]
		}
		maxqueue = append(maxqueue, i)

		// 3. 将窗口中的最大值入结果集
		result = append(result, nums[maxqueue[0]])
	}

	return result
}
