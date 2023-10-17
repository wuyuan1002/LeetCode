package main

// 90. 子集 II

// 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。

// subsetsWithDup .
// 相当于求nums中 0个数字的组合 + 1个数字的组合 + 2个数字的组合 + ... + len(nums)个数字的组合 的并集
func subsetsWithDup(nums []int) [][]int {
	// 因为有重复元素, 所以先排序方便剪枝
	quickSort90(nums, 0, len(nums)-1)

	res := make([]int, 0)      // 一次回溯过程中的结果 -- 回溯路径
	result := make([][]int, 0) // 总结果集

	// 分别求 nums中 0 ～ len(nums) 个数字的组合的并集
	for i := 0; i <= len(nums); i++ {
		dfsSubsetsWithDup(nums, 0, i, &res, &result)
	}

	return result
}

// dfsSubsetsWithDup 回溯遍历选择列表，求出count个数字的组合
// nums: 选择列表
// start: 每次遍历的起始下标 -- 指定当前层的选择范围 -- nums[start: ]
// count: 这次要选取几个数字 -- 终止条件
// res: 一次回溯过程中的结果 -- 回溯路径
// result: 总结果集
func dfsSubsetsWithDup(nums []int, start int, count int, res *[]int, result *[][]int) {
	// 回溯路径已满足条件 -- 将本次回溯的结果记入总结果集后返回
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
		// 剪枝 -- 跳过重复元素, 防止出现重复子集
		if i > start && nums[i] == nums[i-1] {
			continue
		}

		// 在本层固定一个数字加入回溯路径中
		*res = append(*res, nums[i])
		// 进入下一层选择下一个元素, 因为每个数字只能用一次, 所以下一层开始位置为 i+1
		dfsSubsetsWithDup(nums, i+1, count, res, result)
		// 将本层固定的数字移出回溯路径, 重新固定下一个数字
		*res = (*res)[:len(*res)-1]
	}
}

// quickSort90 快排
func quickSort90(nums []int, left, right int) {
	if len(nums) == 0 || left >= right {
		return
	}

	l, r := left, right
	temp := nums[left]

	for l < r {
		for l < r && nums[r] >= temp {
			r--
		}
		for l < r && nums[l] <= temp {
			l++
		}

		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}

	nums[left], nums[l] = nums[l], nums[left]
	quickSort90(nums, left, l-1)
	quickSort90(nums, r+1, right)
}
