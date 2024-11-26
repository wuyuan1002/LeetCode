package main

// 78. 子集

// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

// subsets .
// leetcode 77. 组合
// leetcode 90. 子集 II
//
// 递归回溯
// 相当于求nums中 0个数字的组合 + 1个数字的组合 + 2个数字的组合 + ... + len(nums)个数字的组合 的并集
func subsets(nums []int) [][]int {
	res := make([]int, 0)      // 一次回溯过程中的结果 -- 回溯路径
	result := make([][]int, 0) // 总结果集

	// 分别求 nums中 0 ～ len(nums) 个数字的组合的并集
	for k := 0; k <= len(nums); k++ {
		dfsSubsets(nums, 0, k, &res, &result)
	}

	return result
}

// dfsSubsets 回溯遍历选择列表，求出count个数字的组合
// nums: 选择列表
// start: 每次遍历的起始下标 -- 指定当前层的选择范围 -- nums[start: ]
// count: 这次要选取几个数字 -- 终止条件
// res: 一次回溯过程中的结果 -- 回溯路径
// result: 总结果集
func dfsSubsets(nums []int, start int, count int, res *[]int, result *[][]int) {
	// 回溯路径已满足条件 -- 将本次回溯的结果计入总结果集后返回
	if len(*res) == count {
		temp := make([]int, count)
		copy(temp, *res)
		*result = append(*result, temp)
		return
	}

	// 从本层固定一个值, 然后进入下一层固定下一个值
	for i := start; i < len(nums); i++ {
		// 剪枝 -- 若 已有元素+剩余元素 < count 时可以直接跳出，因为无论如何都不会满足总数等于count了
		if len(*res)+len(nums)-start+1 < count {
			return
		}

		// 在本层固定一个数字加入回溯路径中
		*res = append(*res, nums[i])
		// 进入下一层选择下一个元素, 因为每个数字只能用一次, 所以下一层开始位置为 i+1
		dfsSubsets(nums, i+1, count, res, result)
		// 将本层固定的数字移出回溯路径, 重新固定下一个数字
		*res = (*res)[:len(*res)-1]
	}
}
