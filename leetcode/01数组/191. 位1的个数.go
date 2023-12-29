package main

// 191. 位1的个数

// 编写一个函数，输入是一个无符号整数（以二进制串的形式），
// 返回其二进制表达式中数字位数为 '1' 的个数（也被称为汉明重量）。

// hammingWeight .
// 1. 使用一个无符号整数1不断与数字的每一位进行与运算，如果是1则累加结果，然后左移i，继续验证整数的下一位 -- 数字是几位二进制数字，就需要循环几次，因为循环到最后需要让i溢出变为0 -- 这里需要循环32次
func hammingWeight1(num uint32) int {
	count := 0
	var i uint32 = 1
	for ; i != 0; i = i << 1 {
		if num&i != 0 {
			count++
		}
	}
	return count
}

// hammingWeight2 .
// 2. 把一个整数减1再和原来的整数作与运算，会把该整数最右面的1变成0。那么一个整数中有多少个1就可以做多少次这样的操作
func hammingWeight2(num uint32) int {
	count := 0
	for num != 0 {
		count++
		num &= num - 1
	}
	return count
}
