package main

import (
	"math"
	"strings"
)

// 剑指 Offer 67. 把字符串转换成整数

// 写一个函数 StrToInt，实现把字符串转换成整数这个功能。不能使用 atoi 或者其他类似的库函数。
// 首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。
// 当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。
// 该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。
// 注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。
// 在任何情况下，若函数不能进行有效的转换时，请返回 0。

// 说明：
// 假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为[−231, 231− 1]。如果数值超过这个范围，请返回 INT_MAX (231− 1) 或INT_MIN (−231) 。

// func main() {
// 	fmt.Println(strToInt("9223372036854775808"))
// }

func strToInt(str string) int {

	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}

	sign, i := 1, 0 // 数字是正是负，开数字的开始下标
	// 查看符号位
	if str[i] == '-' {
		sign = -1
		i = 1
	} else if str[i] == '+' {
		i = 1
	}

	var res int
	// 开始处理数值
	for ; i < len(str); i++ {
		// 遇到第一个不符合的 -- 跳出
		if str[i] < '0' || str[i] > '9' {
			break
		}
		// 处理越界的情况
		res = res*10 + int(str[i]-'0')
		if res > math.MaxInt32 {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
	}

	return sign * res
}
