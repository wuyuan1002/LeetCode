package main

// 509. 斐波那契数

// 斐波那契数 （通常用 F(n) 表示）形成的序列称为 斐波那契数列 。
// 该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
//
// F(0) = 0，F(1) = 1
// F(n) = F(n - 1) + F(n - 2)，其中 n > 1
// 给定 n ，请计算 F(n) 。

// fib .
// dp[i]表示大小为i的斐波那契数
// 递推公式：dp[i] = dp[i - 1] + dp[i - 2]
func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	// 构造并初始化dp数组
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	// 开始递推
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	// 返回结果
	return dp[n]
}

// fib1 .
// 从动态规划解法可以发现，对于dp数组，我们其实只使用dp[i-1]和dp[i-2]两个中间结果，
// 并不需要使用dp[i-3]、dp[i-4]等等其它数字，所以我们可以只使用两个变量来保存dp[i-1]和dp[i-2]，
// 而不是使用一个数字保存所有数字的中间结果
func fib1(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	result := 0          // 总结果
	fib_1, fib_2 := 1, 0 // 前一个dp[i-1]和前两个dp[i-2]数字的斐波那契数

	// 开始遍历求斐波那契数，并保存前一个和前两个的中间结果
	for i := 2; i <= n; i++ {
		result = fib_1 + fib_2
		fib_2 = fib_1
		fib_1 = result
	}

	return result
}
