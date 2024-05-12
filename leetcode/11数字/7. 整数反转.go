package main

import "math"

// 7. 整数反转

// 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
//
// 如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
//
// 假设环境不允许存储 64 位整数（有符号或无符号）。

// reverse .
func reverse(x int) int {
	result := 0

	for x != 0 {
		// 获取当前数字的尾部元素
		tail := x % 10

		// 判断后面在计算result时是否会溢出
		if result < math.MinInt32/10 || result > math.MaxInt32/10 {
			// 判断 result*10 是否会溢出 -- 若result已经比 MaxInt32/10 要大了，那么后面 result*10 时就一定会溢出
			return 0
		} else if result == math.MinInt32/10 && tail > 8 {
			// 若当前result正好等于 MinInt32/10，那么新加入的个位数不能大于 -2147483648 的最后一位，否则就会溢出
			return 0
		} else if result == math.MaxInt32/10 && tail > 7 {
			// 若当前result正好等于 MaxInt32/10，那么新加入的个位数不能大于 2147483647 的最后一位，否则就会溢出
			return 0
		}

		// 若不会溢出则将该位反转加入到结果的尾部
		result = result*10 + tail
		x /= 10
	}

	return result
}
