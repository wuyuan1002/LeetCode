package main

import "math"

// 414. 第三大的数

// 给你一个非空数组，返回此数组中 第三大的数 。如果不存在，则返回数组中最大的数。

// thirdMax .
// 1. 使用三个变量分别记录最大、第二大和第三大的数子，一次遍历数组不断更新三个值，遍历结束后返回第三大的数字
// 2. 快排partition -- 快排过程中若发现当前归位了第三大的数字则直接返回不继续排序
func thirdMax(nums []int) int {
	a, b, c := math.MinInt64, math.MinInt64, math.MinInt64
	for _, num := range nums {
		if num > a {
			a, b, c = num, a, b
		} else if num > b && num < a {
			b, c = num, b
		} else if num > c && num < b {
			c = num
		}
	}
	if c == math.MinInt64 {
		return a
	}
	return c
}
