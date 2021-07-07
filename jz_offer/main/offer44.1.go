package main

import (
	"math"
)

// 剑指 Offer 44. 数字序列中某一位的数字

// 数字以0123456789101112131415…的格式序列化到一个字符序列中。
// 在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。
//
// 请写一个函数，求任意第n位对应的数字。

func main() {

}

func findNthDigit1(n int) int {
	if n < 10 {
		return n
	}
	digits := 1
	for {
		nums := count(digits)
		if n > digits*nums {
			n -= digits * nums
			digits++
		} else {
			preCount := n / digits       // 目标数前有多少个digits位数
			index := n - preCount*digits // 第几位
			t := int(math.Pow(10, float64(digits-1))) + index
			return getDigit(t, index, digits)
		}
	}
}

// 统计n位数有多少个,如3位数有900个,1000-100=900
func count(n int) int {
	if n == 1 {
		return 10
	}
	return int(math.Pow(10, float64(n)) - math.Pow(10, float64(n-1)))
}

// 获取整数n的第index位，如871的第0位为8
func getDigit(n, index, amount int) int {
	res := 0
	for index < amount {
		res = n - n/10*10
		n = n / 10
		index++
	}
	return res
}
