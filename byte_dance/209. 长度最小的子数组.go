package main

import (
	"math"
)

// 209. 长度最小的子数组

// 给定一个含有n个正整数的数组和一个正整数 target 。
//
// 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组[numsl, numsl+1, ..., numsr-1, numsr] ，
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

	i, j := 0, 0            // 左右指针
	sum := 0                // 窗口中数字的和
	minLen := math.MaxInt32 // 最小长度

	for j < len(nums) {
		sum += nums[j]
		for sum >= target && i <= j {
			minLen = min(minLen, j-i+1)
			sum -= nums[i]
			i++
		}
		j++
	}

	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}
