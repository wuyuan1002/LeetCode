package main

import (
	"fmt"
)

// 64. 最小路径和

// 给定一个包含非负整数的 m x n 网格 grid ，
// 请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}

// offer 47
// 动态规划
func minPathSum(grid [][]int) int {
	if grid == nil || len(grid) == 0 || grid[0] == nil || len(grid[0]) == 0 {
		return 0
	}

	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}

	products := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if j == 0 { // 第一列
				products[j] = grid[i][j] + products[j]
			} else if i == 0 { // 第一行
				products[j] = grid[i][j] + products[j-1]
			} else {
				products[j] = grid[i][j] + min(products[j], products[j-1])
			}
		}
	}

	return products[len(grid[0])-1]
}
