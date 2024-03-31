package main

import "sort"

// 1387. 将整数按权重排序

// 我们将整数 x 的 权重 定义为按照下述规则将 x 变成 1 所需要的步数：
//
// 如果 x 是偶数，那么 x = x / 2
// 如果 x 是奇数，那么 x = 3 * x + 1
// 比方说，x=3 的权重为 7 。因为 3 需要 7 步变成 1 （3 --> 10 --> 5 --> 16 --> 8 --> 4 --> 2 --> 1）。
//
// 给你三个整数 lo， hi 和 k 。你的任务是将区间 [lo, hi] 之间的整数按照它们的权重 升序排序 ，如果大于等于 2 个整数有 相同 的权重，那么按照数字自身的数值 升序排序 。
//
// 请你返回区间 [lo, hi] 之间的整数按权重排序后的第 k 个数。
//
// 注意，题目保证对于任意整数 x （lo <= x <= hi） ，它变成 1 所需要的步数是一个 32 位有符号整数。

// getKth .
// 对数字进行排序，排序的比较规则不再是直接比较两数字大小，而是使用数字权重进行排序，
// 有两个优化点：
// 1. 为防止相同数字被多次计算权重，所以可以使用动态规划对计算过的数字的权重进行记录 -- dp[num]表示数字num的权重
// 2. 因为只需要返回第k个数字，所以可以自定义快排partition，在排序找到第k个数字后便可停止排序
func getKth(lo int, hi int, k int) int {
	// 构造要排序的数组
	nums := make([]int, 0, hi-lo+1)
	for i := lo; i <= hi; i++ {
		nums = append(nums, i)
	}

	// dp数组 -- dp[num]表示数字num的权重
	dp := make(map[int]int)

	// 对数组进行排序 -- 权重不同时按照权重进行排序，权重相同时按照数字本身进行排序
	sort.Slice(nums, func(i, j int) bool {
		if getWeight(nums[i], dp) != getWeight(nums[j], dp) {
			return getWeight(nums[i], dp) < getWeight(nums[j], dp)
		} else {
			return nums[i] < nums[j]
		}
	})

	// 返回第k个数字
	return nums[k-1]
}

// getWeight 获取数字num的权重
func getWeight(num int, dp map[int]int) int {
	// 若数字的权重已经在dp中，则直接返回
	if weight, ok := dp[num]; ok {
		return weight
	}

	// 计算当前数字的权重
	if num == 1 { // 1的权重为0
		dp[num] = 0
	} else if num%2 == 0 { // 偶数
		dp[num] = getWeight(num/2, dp) + 1
	} else { // 奇数
		dp[num] = getWeight(num*3+1, dp) + 1
	}

	return dp[num]
}
