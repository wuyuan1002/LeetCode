package main

import "math"

// 209. 长度最小的子数组

// 给定一个含有n个正整数的数组和一个正整数 target 。
//
// 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组[numsl, numsl+1, ..., numsr-1, numsr] ，
// 并返回其长度。如果不存在符合条件的子数组，返回 0 。

// minSubArrayLen .
// 同 leetcode 713. 乘积小于 K 的子数组
//
// 双指针, 滑动窗口
// 右指针不断向前移动把数字加到sum, 当发现 sum >= target 时, 开始向前移动左指针, 并不断记录窗口大小
func minSubArrayLen(target int, nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	sum, minLen := 0, math.MaxInt // 当前和、子数组最小长度
	// 定义快慢指针, 不断向前移动右指针
	for l, r := 0, 0; r < len(nums); r++ {
		// 将右指针的值加到当前和上, 当发现sum >= target时开始向前移动左指针并记录最小子数组长度
		for sum += nums[r]; sum >= target && l <= r; l++ {
			minLen = min(minLen, r-l+1)
			sum -= nums[l]
		}
	}

	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}

// min 求最小值
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
