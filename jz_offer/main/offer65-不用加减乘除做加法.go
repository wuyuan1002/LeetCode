package main

// 剑指 Offer 65. 不用加减乘除做加法

// 写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。

// func main() {

// }

func add(a int, b int) int {
	if a == 0 {
		return b
	} else if b == 0 {
		return a
	} else {
		return add(a^b, a&b<<1)
	}
}
