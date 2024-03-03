package main

// 258. 各位相加

// 给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。返回这个结果。

// addDigits .
// 1. 取余得到每一位数字然后相加
// 2. 因为数字num是十进制数字，所以100x+10y+z=99x+9y+x+y+z
func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	if num%9 == 0 {
		return 9
	}
	return num % 9

	// 即：
	// return (num-1)%9 + 1
}
