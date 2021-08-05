package main

import "math"

// 209. 长度最小的子数组

func main() {

}

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
