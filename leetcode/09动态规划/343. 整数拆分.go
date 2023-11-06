package main

// 343. 整数拆分

// 给定一个正整数 n ，将其拆分为 k 个 正整数 的和（ k >= 2 ），并使这些整数的乘积最大化。
//
// 返回 你可以获得的最大乘积 。

// integerBreak .
// 同剪绳子
// 背包问题 -- 最值问题 -- dp[i] = max(dp[i], dp[j] * dp[i-j])
// dp[i]表示拆分数字i所能得到的最大乘积
// dp[i] = max(dp[i], dp[j] * dp[i-j])
func integerBreak(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}

	// 构造dp数组 -- 用于存放拆分 1～n 的每个数字能得到的最大乘积
	dp := make([]int, n+1)

	// 初始化dp数组 -- 最优子结构 -- 在拆分到123后就不要继续拆了
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3

	// 开始拆分 1～n 的每个数字，求它们所能拆分得到的最大乘积
	for i := 4; i <= n; i++ {
		// 不断拆分数字i，计算数字i拆分得到的最大乘积
		for j := 1; j <= i/2; j++ {
			// 将i拆分为j和i-j的最大乘积 == 拆分j的最大乘积 * 拆分i-j的最大乘积
			// 同时与前一次的拆分方式取最大值便为到目前为止拆分i所能得到的最大乘积
			dp[i] = max(dp[i], dp[j]*dp[i-j])
		}
	}

	// 返回dp数组中拆分n的最大乘积
	return dp[n]
}
