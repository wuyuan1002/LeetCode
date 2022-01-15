package main

import (
	"math"
)

// 414. 第三大的数

// 给你一个非空数组，返回此数组中 第三大的数 。如果不存在，则返回数组中最大的数。

// func main() {
// 	fmt.Println(thirdMax([]int{5, 2, 2}))
// }

// 1. 排序
// 2. 快排partition -- 类似215
// 3. 维护一个只存3个元素的数组，遍历nums时将前三大的元素入数组
// 4. 同3，只是使用三个变量维护前3大的元素
func thirdMax(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	// 需注意重复元素 -- 需要排除掉
	// 使用3个变量维护前3大的数字
	a, b, c := math.MinInt64, math.MinInt64, math.MinInt64
	for _, n := range nums {
		if n > a {
			a, b, c = n, a, b
		} else if a > n && n > b {
			b, c = n, b
		} else if b > n && n > c {
			c = n
		}
	}

	// 若不存在第3大的数则返回最大的那个
	if c == math.MinInt64 {
		return a
	}
	return c
}
