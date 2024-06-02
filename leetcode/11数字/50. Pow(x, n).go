package main

// 50. Pow(x, n)

// 实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）。

// myPow .
// 快速幂算法
// 求2的12次方，最简单的办法是将2连乘12次，而更简单的方法是：
// 求2的12次方，相当于先求出2的6次方，然后将两个2的6次方相乘
// 求2的13次方，相当于先求出2的6次方，然后将两个2的6次方相乘再乘一个2
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	} else {
		return 1 / quickMul(x, -n)
	}
}

// quickMul 求x的n次方
func quickMul(x float64, n int) float64 {
	// 任何数的0次方都得1
	if n == 0 {
		return 1
	}

	// 先求出x的n/2次方
	pow := quickMul(x, n/2)

	// 若为偶数次方则为两个得数相乘，若为奇数次方则为两个得数相乘再多乘一次x
	if n%2 == 0 {
		return pow * pow
	} else {
		return pow * pow * x
	}
}
