package main

// 977. 有序数组的平方

// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，
// 要求也按 非递减顺序 排序。

// sortedSquares 双指针
// 负数的平方可能成为最大值，因此，新数组的最大值出现在旧数组两端
func sortedSquares(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	result := make([]int, len(nums))
	l, r := 0, len(nums)-1
	for i := len(result) - 1; l <= r; i-- {
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
