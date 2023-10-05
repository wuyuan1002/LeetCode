package main

import "fmt"

// 239. 滑动窗口最大值

// 给你一个整数数组 nums，有一个大小为k的滑动窗口从数组的最左侧移动到数组的最右侧。
// 你只可以看到在滑动窗口内的 k个数字。滑动窗口每次只向右移动一位。
//
// 返回滑动窗口中的最大值。

// maxSlidingWindow .
// 同 Offer 59、leetcode 155. 最小栈
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

func main() {
	a := []int{-7, -8, 7, 5, 7, 1, 6, 0}
	fmt.Println(maxSlidingWindow(a, 4))
}
