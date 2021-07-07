package main

import (
	"fmt"
)

// 416. 分割等和子集

// 给你一个 只包含正整数 的 非空 数组 nums 。
// 请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

func main() {
	fmt.Println(canPartition([]int{3, 3, 3, 4, 5}))
}

// Hot100 322, 518，39(与回溯法的区别)
// 背包问题 - 存在问题 -- dp[i]=dp[i] || dp[i-num]
func canPartition(nums []int) bool {
	if nums == nil || len(nums) == 0 {
		return true
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		// 总和为奇数，无法分成相等的两部分
		return false
	}

	target := sum / 2 // 每个子数组的目标和 -- 零钱兑换里的总金额

	dp := make([]bool, target+1) // dp[i]:是否存在子集和为i
	dp[0] = true                 // target=0不需要选择任何元素，所以是可以实现的

	for _, num := range nums { // 遍历选择列表 -- 所有零钱
		for i := target; i >= num; i-- { // 遍历目标值 -- 目标值target -- 每个物品只能被使用1次时就应该倒序遍历
			dp[i] = dp[i] || dp[i-num]
		}
	}

	return dp[target]
}
