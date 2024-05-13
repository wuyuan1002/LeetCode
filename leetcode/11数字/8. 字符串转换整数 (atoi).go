package main

import (
	"math"
	"strings"
)

// 8. 字符串转换整数 (atoi)

// 请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数。
//
// 函数 myAtoi(string s) 的算法如下：
//
// 空格：读入字符串并丢弃无用的前导空格（" "）
// 符号：检查下一个字符（假设还未到字符末尾）为 '-' 还是 '+'。如果两者都不存在，则假定结果为正。
// 转换：通过跳过前置零来读取该整数，直到遇到非数字字符或到达字符串的结尾。如果没有读取数字，则结果为0。
// 舍入：如果整数数超过 32 位有符号整数范围 [−231,  231 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231 的整数应该被舍入为 −231 ，大于 231 − 1 的整数应该被舍入为 231 − 1 。
// 返回整数作为最终结果。

// myAtoi .
func myAtoi(s string) int {
	result := 0

	// 去除掉字符串两端空格
	s = strings.TrimSpace(s)

	// 符号位为正还是负
	sign := 1

	for i, c := range s {
		if i == 0 && (c == '-' || c == '+') {
			// 若首位是符号位
			if c == '-' {
				sign = -1
			}
			continue
		} else if c < '0' || c > '9' {
			// 若读到非数字字符则终止解析，直接返回当前结果
			break
		} else if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && c > '7') {
			// 判断结果是否会溢出
			if sign == -1 {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}

		// 将当前数字加大结果末尾
		result = result*10 + int(c-'0')
	}

	return result * sign
}
