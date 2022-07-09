package main

import (
	"math"
)

// 322. 零钱兑换

// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
// 如果没有任何一种硬币组合能组成总金额，返回-1。
//
// 你可以认为每种硬币的数量是无限的。

// func main() {
// 	fmt.Println(coinChange1([]int{5, 3, 1}, 7))
// }

// --------------------

// 背包问题: 给定一个target(直接或间接求出)，给定一个选择列表，求达到target的最优解、解的个数、是否存在等问题

// 背包问题的转移方程:
// 1、最值问题(求最优解): dp[i] = max/min(dp[i], dp[i-num]+1) 或 dp[i] = max/min(dp[i], dp[i-num]+num)
// Hot100 322、1049
// 类似题: 718、1143、279

// 2、存在问题(是否存在): dp[i] = dp[i] || dp[i-num]
// Hot100 416

// 3、组合问题(解的个数): dp[i] += dp[i-num]
// 只是求出解的个数，而不是求出每个解是什么 -- 见Hot100 39
// Hot100 518、39、40、494、377

// --------------------

// 内层循环从小到大还是从大到小遍历(遍历背包容量, 目标值):
// 01背包: 每个物品只能被使用1次，因此，内层循环应该倒序遍历 Hot100 416、494、1049
// 完全背包: 每个物品可以使用多次，因此，内层循环应该正序遍历 Hot100 322、39、518

// 内外层循环能不能颠倒(先遍历物品嵌套遍历背包容量还是先遍历背包容量嵌套遍历物品呢?):
// 若使用二维dp数组, 01背包和完全背包都可以颠倒顺序
// 若使用的是一维dp滚动数组, 01背包必须是先遍历物品嵌套遍历背包容量, 完全背包可以颠倒顺序

// --------------------

// 类似Hot100 518, 416，39(与回溯法的区别)
// 动态规划 -- dp[i]表示i元的最少硬币数 -- dp[i] = min(dp[i], dp[i-num]+1)
func coinChange(coins []int, amount int) int {
	if coins == nil || len(coins) == 0 || amount < 0 {
		return -1
	}

	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}

	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for _, coin := range coins { // 遍历每一个物品
		for i := 1; i <= amount; i++ { // 求选择这个物品后构成目标值的最优解 -- 每个物品只能被使用1次时应该是倒序遍历，如Hot100 416、494
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// ----------
// 第二次做
// 类似Hot100 518, 416，39(与回溯法的区别)
// 动态规划 -- dp[i]表示i元的最少硬币数 -- dp[i] = min(dp[i], dp[i-num]+1)
func coinChange1(coins []int, amount int) int {
	if coins == nil || len(coins) == 0 || amount < 0 {
		return -1
	}

	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}

	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for _, coin := range coins {
		for j := 1; j <= amount; j++ {
			if j >= coin {
				dp[j] = min(dp[j-coin]+1, dp[j])
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
