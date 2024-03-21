package main

// 367. 有效的完全平方数

// 给你一个正整数 num 。如果 num 是一个完全平方数，则返回 true ，否则返回 false 。
//
// 完全平方数 是一个可以写成某个整数的平方的整数。换句话说，它可以写成某个整数和自身的乘积。
//
// 不能使用任何内置的库函数，如  sqrt 。

// isPerfectSquare .
// 1. 二分查找 -- 若结果为result，那么小于result的一侧必然不满足result * result >= num，大于result的一侧则必然满足
// 2. 递减奇数 -- 任何完全平方数都可以是1+3+5+7+...的连续奇数只和，只需将num逐渐递减从1开始的奇数，若能减为0则说明是完全平方数
// 3. 牛顿迭代法
func isPerfectSquare(num int) bool {
	l, r := 0, num
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid < num {
			l = mid + 1
		} else if mid*mid > num {
			r = mid - 1
		} else {
			return true
		}
	}
	return false
}
