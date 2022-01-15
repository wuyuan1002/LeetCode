package main

// 494. 目标和

// 给你一个整数数组 nums 和一个整数 target 。
//
// 向数组中的每个整数前添加'+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
//
// 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，
// 然后串联起来得到表达式 "+2-1" 。
// 返回可以通过上述方法构造的、运算结果等于 target 的不同表达式的数目。

// func main() {
// 	fmt.Println(findTargetSumWays1([]int{1, 1, 1, 1, 1}, 3))
// }

// 1. 回溯法 -- 时间复杂度太高
func findTargetSumWays(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	result := 0
	dfs13(nums, 0, 0, target, &result)
	return result
}

func dfs13(nums []int, index, currentSum, target int, result *int) {
	if index == len(nums) {
		if currentSum == target {
			*result++
		}
		return
	}

	dfs13(nums, index+1, currentSum+nums[index], target, result)
	dfs13(nums, index+1, currentSum-nums[index], target, result)
}

// 2. 动态规划 -- 背包问题 - 组合问题
// 同Hot100 518，类似Hot100 416
func findTargetSumWays1(nums []int, S int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	// left组合 - right组合 = target
	// left + right等于sum，而sum是固定的
	// left - (sum - left) = target -> left = (target + sum)/2

	sum := 0
	for _, v := range nums {
		sum += v
	}
	if S > sum || (sum+S)%2 == 1 {
		return 0
	}
	// 计算目标值 -- 背包大小
	target := (sum + S) / 2
	if target < 0 {
		return 0
	}

	// 定义dp数组
	dp := make([]int, target+1) // dp[i]: 和为i的子集个数
	dp[0] = 1

	for _, num := range nums { // 遍历选择列表 -- 所有零钱
		for i := target; i >= num; i-- { // 遍历目标值 -- target -- 每个物品只能被使用1次时就应该倒序遍历
			dp[i] += dp[i-num]
		}
	}

	return dp[target]
}
