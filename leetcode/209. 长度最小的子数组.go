package main

import "math"

// 209. 长度最小的子数组
// 给定一个含有 n 个正整数的数组和一个正整数 target 。

// 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，
// 并返回其长度。如果不存在符合条件的子数组，返回 0 。

// func main() {

// }

// 滑动窗口
func minSubArrayLen(target int, nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}

	l, r := 0, 0
	sum := 0
	minLen := math.MaxInt32
	for r < len(nums) {
		for sum += nums[r]; sum >= target && l <= r; {
			minLen = min(minLen, r-l+1)
			sum -= nums[l]
			l++
		}
		r++
	}

	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

// 第二次做
func minSubArrayLen1(target int, nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	minLen := math.MaxInt32
	sum := 0
	l, r := 0, 0
	for ; r < len(nums); r++ {
		sum += nums[r]
		for ; sum >= target && l <= r; l++ {
			minLen = min(minLen, r-l+1)
			sum -= nums[l]
		}
	}

	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// 第三次做
func minSubArrayLen2(target int, nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	minLen := math.MaxInt32
	sum := 0
	l, r := 0, 0
	for ; r < len(nums); r++ {
		sum += nums[r]
		for ; sum >= target && l <= r; l++ {
			minLen = min(minLen, r-l+1)
			sum -= nums[l]
		}
	}

	if minLen == math.MaxInt32 {
		return 0
	}

	return minLen
}
