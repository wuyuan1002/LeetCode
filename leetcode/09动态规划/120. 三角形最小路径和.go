package main

import "math"

// 120. 三角形最小路径和

// 给定一个三角形 triangle ，找出自顶向下的最小路径和ÏÏ。
//
// 每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。

// minimumTotal .
// dp[i][j]表示triangle[i, j]位置的最小路径和
// dp[i][j] = triangle[i][j] + min(dp[i-1][j-1], dp[i-1][j])
func minimumTotal(triangle [][]int) int {
	// 构造dp数组 -- dp[i][j]表示triangle[i, j]位置的最小路径和
	dp := make([][]int, len(triangle))
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
	}

	// 初始化dp -- 第一个坐标为三角形顶部，其最小路径和就是本身，而其它位置的最小路径和都依赖其上一行的结果
	dp[0][0] = triangle[0][0]

	// 开始dp -- 求每个位置的最小路径和，从第二行开始求
	for i := 1; i < len(triangle); i++ {
		for j := 0; j <= i; j++ {
			if j == 0 { // 第一列 -- 只能是正上
				dp[i][j] = triangle[i][j] + dp[i-1][j]
			} else if j == i { // 最后一列 -- 只能是左上
				dp[i][j] = triangle[i][j] + dp[i-1][j-1]
			} else { // 其它 -- 由左上和正上得出
				dp[i][j] = triangle[i][j] + min(dp[i-1][j-1], dp[i-1][j])
			}
		}
	}

	// 查找并返回底边（最后一行）的最小路径
	result := math.MaxInt32
	for _, n := range dp[len(triangle)-1] {
		result = min(result, n)
	}
	return result
}
