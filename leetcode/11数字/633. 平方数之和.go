package main

import "math"

// 633. 平方数之和

// 给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c 。

// judgeSquareSum .
// a的平方 + b的平方 = c，因此a和b一定都在[0, 根号c]的范围内
// 因此在[0, 根号c]内使用双指针进行搜索即可
func judgeSquareSum(c int) bool {
	l, r := 0, int(math.Sqrt(float64(c)))

	// a和b可以相等，所以此处要用<= -- 如 2 = 1*1 + 1*1
	for l <= r {
		sum := l*l + r*r
		if sum < c {
			l++
		} else if sum > c {
			r--
		} else {
			return true
		}
	}

	return false
}
