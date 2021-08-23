package main

// 343. 整数拆分

func main() {

}

// 同 剪绳子
func integerBreak(n int) int {
	if n < 2 {
		return 0
	} else if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3

	for i := 4; i <= n; i++ {
		res := 0
		for j := 1; j <= i/2; j++ {
			res = max(res, dp[j]*dp[i-j])
		}
		dp[i] = res
	}

	return dp[n]
}
