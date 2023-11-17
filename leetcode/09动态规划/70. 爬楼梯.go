package main

// 70. 爬楼梯

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

// climbStairs .
// dp[i]表示爬到第i阶台阶有几种跳法
// 因为一次可以跳1个或2个台阶，所以对于第i级台阶，可以是从前一级跳一个台阶跳上来，也可以是从前两级跳两个台阶跳上来，所以第i级台阶的跳法总数为前一级的跳法总数+前两级的跳法总数
// 递推公式：dp[i] = dp[i - 1] + dp[i - 2]
func climbStairs(n int) int {
	if n < 2 {
		return 1
	}

	// 构造并初始化dp数组
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	// 开始递推
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	// 返回结果
	return dp[n]
}

// climbStairs1 .
// 从dp数组可以发现，我们其实只使用dp[i-1]和dp[i-2]两个中间结果，
// 并不需要使用dp[i-3]、dp[i-4]等等其它数字，所以我们可以只使用两个变量来保存dp[i-1]和dp[i-2]，
// 而不是使用一个数字保存所有数字的中间结果
func climbStairs1(n int) int {
	if n < 2 {
		return 1
	}

	pre1, pre2 := 1, 1 // 前一级和前两级跳法总数
	result := 0        // 当前级跳法总数

	// 开始遍历求每一级的跳法总数，并保存前一个和前两个的中间结果
	for i := 2; i <= n; i++ {
		result = pre1 + pre2
		pre2 = pre1
		pre1 = result
	}

	return result
}
