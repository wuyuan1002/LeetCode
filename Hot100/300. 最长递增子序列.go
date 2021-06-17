package main

import (
	"fmt"
)

// 300. 最长递增子序列

// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
// 子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。
// 例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

func main() {
	// fmt.Println(lengthOfLIS([]int{4, 10, 4, 3, 8, 9}))
	fmt.Println(lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
}

// 类似 Hot100 128
// 动态规划 -- dp[i]表示以第i个数字结尾的最长递增序列的长度 -- dp[i] = max(dp[i], dp[j] + 1) for j in [0, i)
func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	products := make([]int, len(nums))
	for i := range products {
		products[i] = 1
	}

	res := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				products[i] = max(products[i], products[j]+1)
			}
		}
		res = max(res, products[i])
	}

	return res
}
