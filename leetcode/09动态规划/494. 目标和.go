package main

// 494. 目标和

// 给你一个非负整数数组 nums 和一个整数 target 。
//
// 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
//
// 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
// 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

// findTargetSumWays .
// 将选择列表中的数字分成两组（left组和right组），使得两组数字的差正好等于target
// 那么整体的和等于左右两组的和：sum = left + right -> right = sum - left
// 左组与右组的差为目标值target：left - right = target
// 而sum和target都是已知的：left - (sum - left) = target -> left = (target + sum) / 2
// 所以最终变成了从选择列表中选择一些数字使得他们的总和等于left有几种选择方式 -> 组合总和 leetcode 377
// left组的和便变成了新的target，我们要求的就是从选择列表中选择数字使得其和为left
//
// 1. 回溯法 -- 可以求出每种组合具体是什么
// 2. 动态规划 -- 01背包 -- 装满容量为left的背包，有几种方法 -- 只能求出有几种组合数，但不知道每种组合是什么
// dp[i]表示选择数字使得其和为i的组合个数 -- 填满容量为i的背包有几种组合情况
// dp[i] = dp[i] + dp[i - num]
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if target > sum || (sum+target)%2 == 1 {
		// sum = 3，target = 8 是无解的
		// sum = 5，target = 2 是无解的
		return 0
	}

	// 获取新的target -- left就是新的target，即背包容量
	left := (target + sum) / 2
	if left < 0 {
		// num = 100, target = -200 是无解的
		return 0
	}

	// 构建dp数组
	dp := make([]int, left+1)

	// 初始化dp数组 -- 其实这个值没有意义，只是如果dp[0] = 0 的话，后面所有推导出来的值都是0了
	dp[0] = 1

	for _, num := range nums { // 遍历选择列表 -- 不断选择数字
		for i := left; i >= num; i-- { // 倒序遍历背包容量，计算并更新填满该容量的背包的组合个数 -- 因为每个数字只能使用一次，所以需要倒序变量背包容量
			dp[i] += dp[i-num]
		}
	}

	// 返回选择数字使得其和为left有几种组合情况
	return dp[left]
}
