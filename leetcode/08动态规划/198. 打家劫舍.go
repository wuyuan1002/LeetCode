package main

// 198. 打家劫舍

// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，
// 影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
// 如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//
// 给定一个代表每个房屋存放金额的非负整数数组，
// 计算你不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

// rob .
// 同 leetcode 70、509
// 偷到某个房屋的最大价值 = max(当前房屋的金额 + 偷到前两个房屋所能得到的最高金额, 偷到前一个房屋所能得到的最高金额)
// dp[i]表示偷到第i间房间所能得到的最高总金额
// dp[i] = max(nums[i] + dp[i-2], dp[i-1])
func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	// 构造dp数组 -- dp[i]表示偷到第i间房间所能偷到的最高金额
	dp := make([]int, len(nums))

	// 初始化dp数组 -- 偷到第1个房间所能得到的最大金额就是房间本身的金额、偷到第2个房间所能得到的最大金额是前两个房间的最大值（因为两个房间只能偷一个，两个都偷会报警）
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	// 开始dp -- 求偷到每个房间所能得到的最大金额
	for i := 2; i < len(nums); i++ {
		dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	}

	// 返回偷到最后一个房间所能得到的最大金额
	return dp[len(nums)-1]
}

// rob1 .
// 从dp数组可以发现，我们其实只使用dp[i-1]和dp[i-2]两个中间结果，
// 并不需要使用dp[i-3]、dp[i-4]等等其它数字，所以我们可以只使用两个变量来保存dp[i-1]和dp[i-2]，
// 而不是使用一个数字保存所有数字的中间结果
func rob1(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	// 偷到某个房屋的最大价值 = max(当前房屋的金额 + 偷到前两个房屋所能得到的最高金额, 偷到前一个房屋所能得到的最高金额)
	pre2 := nums[0]               // 前两个房屋的最大价值
	pre1 := max(nums[0], nums[1]) // 前一个房屋的最大价值

	for i := 2; i < len(nums); i++ {
		// 偷到当前房屋所能得到的最大金额
		price := max(nums[i]+pre2, pre1)

		pre2 = pre1
		pre1 = price
	}

	return pre1
}
