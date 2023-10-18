package main

// 46. 全排列

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

// permute .
// 递归+回溯, 在一层固定一个数字，然后进入下一层固定下一个数字
// 可以使用res回溯列表来存储已经选择好的数字, 也可以使用在原数组中直接交换的方式, 将选好的数字交换到前面来
func permute(nums []int) [][]int {
	result := make([][]int, 0) // 总结果集

	dfsPermute(nums, 0, &result)
	return result
}

// dfsPermute 回溯遍历选择列表，求出nums中数字的全排列
// 可以使用res来存储已经选择好的数字
// 也可以使用直接在原数组中交换数字的方式实现选择数字
//
// nums: 选择列表
// start: 每次遍历的起始下标 -- 指定当前层的选择范围 -- nums[start: ]
// result: 总结果集
func dfsPermute(nums []int, start int, result *[][]int) {
	// 若nums中所有数字都被选择完了, 说明找到了一种排列方式, 将结果记录到总结果中
	if start == len(nums)-1 {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*result = append(*result, temp)
		return
	}

	// 由于本题nums中不包含重复数字, 所以可以不用进行去重剪枝, 而47题就需要

	// 不断从本层选取一个数字交换到最前面, 然后进入下一层选择下一个数字
	for i := start; i < len(nums); i++ {
		// 交换当前数字到最前面表示选择当前数字已被选择
		nums[start], nums[i] = nums[i], nums[start]
		// 进入下一层选择下一个数字
		// 当前层选择的数字被交换到start下标处, 下一层开始位置下标为 start+1
		dfsPermute(nums, start+1, result)
		// 选择当前数字递归结束, 交换回原位置, 继续选择下一个数字进行递归
		nums[start], nums[i] = nums[i], nums[start]
	}
}
