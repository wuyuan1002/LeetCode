package main

import (
	"fmt"
	"math"
)

// 7. 整数反转

// 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
// 如果反转后整数超过 32 位的有符号整数的范围, 就返回0。
//
// 假设环境不允许存储 64 位整数（有符号或无符号）。

func main() {
	fmt.Println(reverse(1534236469))
}

// 1. 转换成字符串，再翻转字符串
// 2. 从末尾挨个取数字
func reverse(x int) int {

	neg := false // 是否为负数
	if x < 0 {
		neg = true
		x = -x
	}

	result := 0 // 总结果
	n := 0      // 最后一位数字
	for x != 0 {
		n = x % 10
		x /= 10

		result = result*10 + n
		if result > math.MaxInt32 {
			// 若超出范围
			return 0
		}
	}

	if neg {
		result = -result
	}

	return result
}
