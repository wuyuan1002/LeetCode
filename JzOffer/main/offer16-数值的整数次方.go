package main

import "fmt"

// 剑指 Offer 16. 数值的整数次方

// 实现函数double Power(double base, int exponent)，求base的exponent次方。
// 不得使用库函数，同时不需要考虑大数问题。

func main() {
	fmt.Printf("%f", myPow(0.00001, 2147483647))
}

func myPow(x float64, n int) float64 {

	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	var un = n
	if n < 0 { // 查看乘方的数是不是负数
		un = -n
	}

	result := myPow(x, un>>1)
	result *= result
	if un&0x1 == 1 { // 若乘方是奇数
		result *= x
	}

	if n < 0 {
		result = 1 / result
	}
	return result
}

// 动态规划
func myPow1(x float64, n int) float64 {
	if x == 0.0 { // 0的任何次方都得0
		return 0.0
	}
	if n == 0 { // 任何数的零次方都得1
		return 1
	}
	var un = n
	if n < 0 { // 查看乘方的数是不是负数
		un = -n
	}

	results := make([]float64, un+1)
	results[0] = 1 // 最优子结构
	results[1] = x

	for i := 2; i <= un; i++ {
		j := i / 2
		if i%2 == 0 {
			results[i] = results[j] * results[j]
		} else {
			results[i] = results[j] * results[i-j]
		}
	}
	result := results[un]
	if n < 0 {
		result = 1 / result
	}
	return result
}
