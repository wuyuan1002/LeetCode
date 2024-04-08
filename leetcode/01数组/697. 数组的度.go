package main

// 697. 数组的度

// 给定一个非空且只包含非负数的整数数组 nums，数组的 度 的定义是指数组里任一元素出现频数的最大值。
// 你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。

// findShortestSubArray .
// 即求数组中能够将出现次数最多的数字全部包含的子数组最短长度
// 1. 滑动窗口
// 2. 求出数组中出现次数最多的数字（可能有多个），比较其首次出现和最后一次出现的下标差，取最小值即为答案
func findShortestSubArray(nums []int) int {
	// 定义一个map，用来存每个数字出现的次数、首次出现的下标、最后一次出现的下标
	type entry struct{ count, left, right int }
	hash := make(map[int]entry)

	// 遍历数组，统计每个数字的信息
	for i, n := range nums {
		if info, ok := hash[n]; ok {
			info.count++
			info.right = i
			hash[n] = info
		} else {
			hash[n] = entry{1, i, i}
		}
	}

	// 总结果、数字最多的出现次数
	result, maxCnt := 0, 0

	// 遍历map，求出出现次数最多的数字的左右下标差值的最小值
	for _, info := range hash {
		if info.count > maxCnt {
			maxCnt = info.count
			result = info.right - info.left + 1
		} else if info.count == maxCnt {
			result = min(result, info.right-info.left+1)
		}
	}

	return result
}
