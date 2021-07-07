package main

import (
	"fmt"
)

// 剑指 Offer 47. 礼物的最大价值

// 在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。
// 你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。
// 给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？

func main() {
	grid := [][]int{{1, 2, 5}, {3, 2, 1}}
	fmt.Println(maxValue(grid))
}

// maxValueDg 递归解法 -- 超出时间限制
func maxValueDg(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	return maxV(grid, 0, 0)
}

func maxV(grid [][]int, i int, j int) int {
	if i >= len(grid) || j >= len(grid[0]) {
		return 0
	}

	rightV := grid[i][j] + maxV(grid, i, j+1)
	downV := grid[i][j] + maxV(grid, i+1, j)
	if rightV > downV {
		return rightV
	}
	return downV
}

// 动态规划解法 -- 我当前位置的最大价值取决于我的上面和左面位置的最大价值
func maxValue(grid [][]int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 动态规划，用长度为列数的一维滚动数组代替二位dp数组
	products := make([]int, len(grid[0]))
	products[0] = grid[0][0]
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i-1 >= 0 && j-1 >= 0 {
				products[j] = grid[i][j] + max(products[j], products[j-1])
			} else if i-1 < 0 && j-1 >= 0 {
				// 第一行的
				products[j] = grid[i][j] + products[j-1]
			} else if i-1 >= 0 && j-1 < 0 {
				// 第一列的
				products[j] = grid[i][j] + products[j]
			}
		}
	}
	return products[len(grid[0])-1]
}

// 第二次做
// 动态规划 -- 我当前位置的最大价值取决于我的上面和左面位置的最大价值
func maxValue1(grid [][]int) int {
	if grid == nil || len(grid) == 0 || grid[0] == nil || len(grid[0]) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	products := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if j == 0 { // 第一列 -- j-1<0会越界异常，特殊处理
				products[j] = grid[i][j] + products[j]
			} else { // 其他列正常计算
				products[j] = grid[i][j] + max(products[j], products[j-1])
			}
		}
	}

	return products[len(products)-1]
}
