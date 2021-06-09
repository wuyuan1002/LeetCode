package main

import (
	"fmt"
)

// 给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的二叉搜索树有多少种？
// 返回满足题意的二叉搜索树的种数。

func main() {
	fmt.Println(numTrees(3))
}

// 动态规划 -- 类似于剪绳子 offer 14
func numTrees(n int) int {

	products := make([]int, n+1)
	products[0] = 1
	products[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			products[i] += products[j-1] * products[i-j]
		}
	}
	return products[n]
}

// 超出时间限制 -- 有重复计算，所以使用上面的动态规划优化
func numTrees1(n int) int {
	if n <= 1 {
		return 1
	}

	num := 0
	for i := 1; i <= n; i++ {
		num += numTrees(i-1) * numTrees(n-i)
	}
	return num
}
