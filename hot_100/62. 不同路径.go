package main

import (
	"fmt"
)

// 62. 不同路径

// 一个机器人位于一个 m x n网格的左上角 （起始点在下图中标记为 “Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
// 问总共有多少条不同的路径？

func main() {
	fmt.Println(uniquePaths1(3, 7))
}

// 回溯法 -- 超时
func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	// 构建地图
	mapp := make([][]int, n)
	for i := 0; i < n; i++ {
		mapp[i] = make([]int, m)
	}

	num := 0
	move(mapp, 0, 0, m, n, &num)

	return num
}

func move(mapp [][]int, i, j, m, n int, num *int) {
	// 若已经到最后一个位置，总结果+1
	if i == m-1 && j == n-1 {
		*num += 1
		return
	}

	// 超出范围直接返回
	if i >= m || j >= n {
		return
	}

	// 向下向右移动
	move(mapp, i+1, j, m, n, num)
	move(mapp, i, j+1, m, n, num)
}

// 类似于offer 47
// 动态规划 -- 到达某一个位置的走法 == 它上面位置的走法个数 + 左面位置的走法个数
func uniquePaths1(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	products := make([]int, n)
	products[0] = 1 // 第一个位置的走法个数为1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 { // 第一列
				products[j] = products[j]
			} else { // 其他
				products[j] = products[j] + products[j-1]
			}
		}
	}
	return products[n-1]
}
