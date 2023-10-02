package main

import (
	"math"
)

// 120. 三角形最小路径和

// 给定一个三角形 triangle ，找出自顶向下的最小路径和。
//
// 每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标
// 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，
// 那么下一步可以移动到下一行的下标 i 或 i + 1 。

// func main() {
// 	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
// 	fmt.Println(minimumTotal(triangle))
// }

// 类似64、118
// 动态规划(二维dp) -- dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
// dp[i][j]表示从三角形顶部走到位置 (i, j) 的最小路径和
func minimumTotal(triangle [][]int) int {
	if triangle == nil || len(triangle) == 0 {
		return 0
	}

	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}

	dp := make([][]int, len(triangle))
	for i := range dp {
		dp[i] = make([]int, i+1)
	}

	// 遍历数组
	dp[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j] // 第一列 -- 只能是上一个
			} else if j == i {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j] // 最后一列 -- 只能在左上
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			}
		}
	}

	// 查找底边的最小路径
	res := math.MaxInt32
	for _, n := range dp[len(triangle)-1] {
		res = min(res, n)
	}
	return res
}

// 一维滚动dp
func minimumTotal1(triangle [][]int) int {
	if triangle == nil || len(triangle) == 0 {
		return 0
	}

	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}

	dp := make([]int, len(triangle))
	dp[0] = triangle[0][0]

	// 遍历数组
	for i := 1; i < len(triangle); i++ {
		for j := i; j >= 0; j-- {
			if j == 0 { // 第一列
				dp[j] = dp[j] + triangle[i][j]
			} else if j == i { // 最后一列
				dp[j] = dp[j-1] + triangle[i][j]
			} else {
				dp[j] = min(dp[j-1], dp[j]) + triangle[i][j]
			}
		}
	}

	// 查找底边的最小路径
	res := math.MaxInt32
	for _, n := range dp {
		res = min(res, n)
	}
	return res
}
