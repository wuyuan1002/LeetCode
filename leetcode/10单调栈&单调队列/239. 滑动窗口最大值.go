package main

// 239. 滑动窗口最大值

// 给你一个整数数组 nums，有一个大小为k的滑动窗口从数组的最左侧移动到数组的最右侧。
// 你只可以看到在滑动窗口内的 k个数字。滑动窗口每次只向右移动一位。
//
// 返回滑动窗口中的最大值。

// ---------
// 什么时候用单调栈什么时候用单调队列呢？
// 求队列的最大最小值时用单调队列：Offer 59-2、leetcode 239. 滑动窗口最大值
// 求栈的最大最小值时用单调栈：leetcode 155. 最小栈
//
// 求最大值时用单调递减，求最小值时用单调递增
// 单调递增栈：从栈顶到栈底是递增的、单调递增队列：从队头到队尾是递增的
// eg：
// 单调递减队列入队时：从单调队列队尾入队，若队尾元素小于被添加元素则将其删掉，直到将被添加元素加入到队尾
// 单调递减队列出队时：若出队元素正好是最大值，则同步从单调队列中出队，即更新队列最大值
// 获取队列最大值：单调递减队列队头元素为当前队列最大值
//
// 单调递增栈入栈时：若添加的元素小于等于单调栈的栈顶元素，则同步入栈到单调栈中，即更新栈内的最小值
// 单调递增栈出栈时：若弹出的元素正好是最小值，则同步从单调栈中弹出，即更新栈内最小值
// 获取栈的最小值：单调递增栈的栈顶元素为当前栈的最小值
// ---------

// maxSlidingWindow .
// 同 Offer 59-2、leetcode 155. 最小栈
// 本体的滑动窗口实际就是队列，向后移动窗口就相当于出队入队同时进行
// 单调递减队列 -- 第一个数是整个队列的最大值，第二个数是单调队列第一个数后面队列中数据的最大值，第三个数是单调队列第二个数后面队列中数据的最大值，以此类推
// 单调递减队列添加元素v时，始终是将队列中所有小于v的数据删除掉，然后将v添加到队列尾部
func maxSlidingWindow(nums []int, k int) []int {
	if nums == nil || len(nums) == 0 || k <= 0 || k > len(nums) {
		return nil
	}

	result := make([]int, 0)   // 结果集
	maxqueue := make([]int, 0) // 单调递减队列

	i := 0
	// 1. 移动i形成第一个窗口，同时构造单调队列
	for ; i < k; i++ {
		// 要添加的元素为nums[i]，从单调队列尾部开始遍历，若队列尾部小于nums[i]则将队尾元素删除，直到队列中的元素全部大于等于nums[i]
		for len(maxqueue) > 0 && nums[i] > maxqueue[len(maxqueue)-1] {
			maxqueue = maxqueue[:len(maxqueue)-1]
		}
		// 将nums[i]添加到队尾
		maxqueue = append(maxqueue, nums[i])
	}

	// 2. 将第一个窗口的最大值入结果集
	result = append(result, maxqueue[0])

	// 3. 向后移动窗口同时记录最大值
	for ; i < len(nums); i++ {
		// 处理出窗口的值 -- 若出窗口的值正好是单调队列的最大值，则先将它移出单调队列
		if nums[i-k] == maxqueue[0] {
			maxqueue = maxqueue[1:]
		}

		// 处理入窗口的值 -- 将入窗口的值加入到单调队列中
		for len(maxqueue) > 0 && nums[i] > maxqueue[len(maxqueue)-1] {
			maxqueue = maxqueue[:len(maxqueue)-1]
		}
		maxqueue = append(maxqueue, nums[i])

		// 将新窗口的最大值加入到结果集
		result = append(result, maxqueue[0])
	}

	return result
}
