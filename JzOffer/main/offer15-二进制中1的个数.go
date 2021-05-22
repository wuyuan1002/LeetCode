package main

import "fmt"

// 剑指 Offer 15. 二进制中1的个数

// 请实现一个函数，输入一个整数（以二进制串形式），输出该数二进制表示中 1 的个数。
// 例如，把 9表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2。
func main() {
	fmt.Printf("%d", hammingWeight3(11))
}

// 数字是几位二进制数字，就需要循环多少次，因为循环到最后需要让i溢出变为0 -- 这里需要循环32次
func hammingWeight(num uint32) int {
	count := 0
	var i uint32 = 1
	for ; i != 0; i = i << 1 {
		if num&i != 0 {
			count++
		}
	}
	return count
}

func hammingWeight1(num uint32) int {
	count := 0
	for num != 0 {
		count++
		num = (num - 1) & num
	}
	return count
}

// 第二次做
func hammingWeight2(num uint32) int {
	count := 0
	var i uint32 = 1
	for ; i != 0; i <<= 1 {
		if num&i != 0 {
			count++
		}
	}
	return count
}

func hammingWeight3(num uint32) int {
	// 把一个整数减1再和原来的整数作与运算，会把该整数最右面的1变成0。那么一个整数中有多少个1就可以做多少次这样的操作
	count := 0
	for num != 0 {
		count++
		num &= num - 1
	}
	return count
}
