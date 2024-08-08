package main

// 977. 有序数组的平方

// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，
// 要求也按 非递减顺序 排序。

// sortedSquares .
// 双指针
// 负数的平方也可能成为最大值，所以新数组的最大值出现在旧数组的两端数字的平方
// 使用双指针从首尾分别选择两个数字，根据绝对值判断两个数字谁在前谁在后
func sortedSquares(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	result := make([]int, len(nums)) // 总结果
	l, r := 0, len(nums)-1           // 左右指针

	// 倒序遍历结果数组，根据旧数组首尾数字的绝对值由大到小填充结果数组
	for i := len(result) - 1; l <= r; i-- {
		// 若左指针的绝对值小于右指针，则说明右指针指向数字的平方更大，填充右指针数字平方进结果数组，否则填充左指针数字平方进结果数组
		if abs(nums[l]) < abs(nums[r]) {
			result[i] = nums[r] * nums[r]
			r--
		} else {
			result[i] = nums[l] * nums[l]
			l++
		}
	}

	return result
}

// abs 求绝对值
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
