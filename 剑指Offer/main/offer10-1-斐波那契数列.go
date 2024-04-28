package main

// 斐波那契数列

// 写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：
//
// F(0) = 0, F(1) = 1
// F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
// 斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。
//
// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

// func main() {
// 	fmt.Printf("%d", fib(8))
// }

const constant = 1000000007

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	fib1 := 0                 // 第n项的前两项
	fib2 := 1                 // 第n项的前一项
	var num int               // 第n项
	for i := 2; i <= n; i++ { // 从第2项开始算
		num = (fib1 + fib2) % constant
		fib1 = fib2
		fib2 = num
	}
	return num
}

// fib1 求斐波那契数列的第n项 -- 递归解法，太慢了，超出时间限制
func fib1(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib1(n-1) + fib1(n-2)
}

// 第二次做 -- 动态规划，后一项等于前两项的和，后面的值取决于前面的值
func fib2(n int) int {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	fib0 := 0
	fib1 := 1

	fib := 0
	for i := 2; i <= n; i++ {
		fib = (fib0 + fib1) % constant
		fib0 = fib1
		fib1 = fib
	}
	return fib
}
