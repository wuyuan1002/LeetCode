package main

// 740. 删除并获得点数

// 给你一个整数数组 nums ，你可以对它进行一些操作。
// 每次操作中，选择任意一个 nums[i] ，删除它并获得 nums[i] 的点数。
// 之后，你必须删除 所有 等于 nums[i] - 1 和 nums[i] + 1 的元素。
//
// 开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。

// deleteAndEarn .
// 同 leetcode 198. 打家劫舍
//
// 当选择元素x后，数组中x-1、x、x+1三个数字将被删除，同时获取x的点数，也就是说选了x后将不能再选择x-1和x+1，
// 因此，可以使用一个数组sums记录数组中所有数字及其出现的和，这就是选择它时所能得到的最大点数，因此，若选择了x，
// 则可以获取sum[x]的点数，且无法再选择x−1和x+1，这就和打家劫舍是一样的逻辑了
//
// dp[i]表示偷到第i个数字所能得到的最高总点数
// dp[i] = max(sum[i] + dp[i-2], dp[i-1])
func deleteAndEarn(nums []int) int {
	// 获取数组中的最大值 -- 最终要偷盗的数组长度必须是maxVal+1（下标0无意义），
	// 因为在被偷盗的数组中，使用下标表示要被偷的数字，所以选择不偷x−1和x+1时是按照下标选择的
	maxNum := 0
	for _, val := range nums {
		maxNum = max(maxNum, val)
	}

	// 构造最终要被偷盗的数组，下标表示数字，其值表示该数字所能贡献的点数 -- 统计每个数字及其出现的和
	sums := make([]int, maxNum+1)
	for _, n := range nums {
		sums[n] += n
	}

	// 求到每个数字所能贡献的点数后，对每个数字进行打家劫舍，分别求偷到每个数字时所能得到的最高总点数
	return rob(sums)
}
