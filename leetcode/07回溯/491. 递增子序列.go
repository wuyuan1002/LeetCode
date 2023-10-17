package main

// 491. 递增子序列

// 给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。
// 你可以按 任意顺序 返回答案。
//
// 数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。

// findSubsequences .
// 同 leetcode 90. 子集 II, 区别在于本题只求递增的子集 --
// 如[4,4,3,2,1]的递增子集只有[4,4], 而子集却包括[[],[2],[3],[4],[2,3],[2,4],[3,4],[4,4],[2,3,4],[2,4,4],[3,4,4],[2,3,4,4]]
func findSubsequences(nums []int) [][]int {
	// 由于只求在原数组顺序中的递增子集, 所以要保证数组顺序,
	// 所以不能通过对数组排序来去重, 而是在每层使用map或set来判断数字是否已经访问过
	if nums == nil || len(nums) < 2 {
		return nil
	}

	res := make([]int, 0)      // 一次回溯过程中的结果 -- 回溯路径
	result := make([][]int, 0) // 总结果集

	// 因为只要递增的子集, 所以集合中元素至少为2个
	// 分别求 nums中 2 ～ len(nums) 个数字的组合的并集
	for i := 2; i <= len(nums); i++ {
		dfsFindSubsequences(nums, 0, i, &res, &result)
	}

	return result
}

// dfsFindSubsequences 回溯遍历选择列表，求出count个数字的组合
// nums: 选择列表
// start: 每次遍历的起始下标 -- 指定当前层的选择范围 -- nums[start: ]
// count: 这次要选取几个数字 -- 终止条件
// res: 一次回溯过程中的结果 -- 回溯路径
// result: 总结果集
func dfsFindSubsequences(nums []int, start int, count int, res *[]int, result *[][]int) {
	// 回溯路径已满足条件 -- 将本次回溯的结果记入总结果集后返回
	if len(*res) == count {
		temp := make([]int, count)
		copy(temp, *res)
		*result = append(*result, temp)
		return
	}

	// 使用set记录本层已经访问过的数字 -- 由于数组未排序，所以使用set去重，而不是像90那样直接比较前一个数字
	visited := make(map[int]struct{})

	// 不断从本层选择满足递增的数字加入到回溯路径中然后进入下一层选择下一个数字
	for i := start; i < len(nums); i++ {
		// 剪枝 -- 只选择满足递增的数字加入到回溯路径
		if len(*res) > 0 && nums[i] < (*res)[len(*res)-1] {
			continue
		}

		// 剪枝 -- 去重, 若当前数字在本层已经被选择过, 则跳过当前数字, 防止出现重复解
		if _, ok := visited[nums[i]]; ok {
			continue
		}

		// 标记当前数字在本层已被选择过
		visited[nums[i]] = struct{}{}

		// 在本层固定一个数字加入回溯路径中
		*res = append(*res, nums[i])
		// 进入下一层选择下一个元素, 因为每个数字只能用一次, 所以下一层开始位置为 i+1
		dfsFindSubsequences(nums, i+1, count, res, result)
		// 将本层固定的数字移出回溯路径, 重新固定下一个数字
		*res = (*res)[:len(*res)-1]
	}
}
