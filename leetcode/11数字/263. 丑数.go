package main

// 263. 丑数

// 丑数 就是只包含质因数 2、3 和 5 的正整数。
// 给你一个整数 n ，请你判断 n 是否为 丑数 。如果是，返回 true ；否则，返回 false 。

// isUgly .
// 将数字不断除以2、3、5，最终得数若为1说明是丑数
// 同 Offer 49. 丑数 -- 求第n个丑数 -- 动态规划，当前丑数等于其前一个丑数x2、x3、x5得到
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	for _, factor := range []int{2, 3, 5} {
		for n%factor == 0 {
			n /= factor
		}
	}
	return n == 1
}
