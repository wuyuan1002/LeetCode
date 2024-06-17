package main

import "math"

// 2786. 访问数组中的位置使分数最大

// 给你一个下标从 0 开始的整数数组 nums 和一个正整数 x 。
//
// 你 一开始 在数组的位置 0 处，你可以按照下述规则访问数组中的其他位置：
//
// 如果你当前在位置 i ，那么你可以移动到满足 i < j 的 任意 位置 j 。
// 对于你访问的位置 i ，你可以获得分数 nums[i] 。
// 如果你从位置 i 移动到位置 j 且 nums[i] 和 nums[j] 的 奇偶性 不同，那么你将失去分数 x 。
// 请你返回你能得到的 最大 得分之和。
//
// 注意 ，你一开始的分数为 nums[0] 。

// maxScore .
// 同 leetcode 70. 爬楼梯
// 只是爬楼梯是每个位置只能由其前一个和前两个位置移动过来，而本题是每个位置都可以由其前面的任意位置移动过来
//
// 构造一个长度为2的dp数组，分别用来存当前位置之前的偶数和奇数的最大得分
// dp[0]表示位置i之前的偶数最大得分
// dp[1]表示位置i之前的奇数最大得分
//
// 因此
// - 若nums[i]为偶数：到达位置i的最大得分 = 位置i的分数 + max(位置i之前的偶数最大得分, 位置i之前的奇数最大得分 - x)
// - 若nums[i]为奇数：到达位置i的最大得分 = 位置i的分数 + max(位置i之前的奇数最大得分, 位置i之前的偶数最大得分 - x)
func maxScore(nums []int, x int) int64 {
	// 总结果 -- 至少是刚开始位置的分数
	result := nums[0]

	// 构造dp数组 -- dp[0]表示位置i之前的偶数最大得分、dp[1]表示位置i之前的奇数最大得分
	dp := [2]int{math.MinInt32, math.MinInt32}

	// 初始化dp数组 -- 更新初始位置的奇偶性分数
	dp[nums[0]%2] = nums[0]

	// 从第二个数字开始遍历每一个数字，计算移动到它所能得到的最大得分，过程中记录总结果
	for i := 1; i < len(nums); i++ {
		parity := nums[i] % 2

		// 计算到达位置i所能得到的的最大得分
		current := max(dp[parity]+nums[i], dp[1-parity]-x+nums[i])
		// 更新dp数组中当前奇偶性的最大得分
		dp[parity] = max(dp[parity], current)

		// 更新总结果
		result = max(result, current)
	}

	return int64(result)
}
