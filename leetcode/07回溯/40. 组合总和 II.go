package main

import "sort"

// 40. 组合总和 II

// 给定一个数组candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。
// candidates中的每个数字在每个组合中只能使用一次。
//
// 注意：解集不能包含重复的组合。

// combinationSum40
// leetcode 39. 组合总和
//
// 递归回溯
// 由于选择列表中的数字存在重复，所以在数字选择过程中需要进行去重减枝
// 每次选定一个数字后进入下一层继续选择下一个数字到回溯路径，直到总和达到目标值说明找到一个解
// 记录结果后向前回溯，将回溯路径的数字不断重新选择，寻找其他的解并进行记录
//
// 此题也类似动态规划01背包问题 -- leetcode 416. 分割等和子集
// 不过动态规划只能求出解的个数，而无法求出每个解是什么，且此题有重复元素，计算时需要去重
func combinationSum40(candidates []int, target int) [][]int {
	if candidates == nil || len(candidates) == 0 {
		return nil
	}

	// 先排序方便剪枝 -- 也可在dfs时使用visited记录已访问的数字来进行去重剪枝
	sort.Ints(candidates)

	res := make([]int, 0)      // 一次回溯过程中的结果 -- 回溯路径
	result := make([][]int, 0) // 总结果集

	dfscombinationSum40(candidates, 0, 0, target, &res, &result)

	return result
}

// dfscombinationSum40 回溯遍历选择列表，记录满足条件的结果
// candidates: 选择列表
// start: 每次遍历的起始下标 -- 指定当前层的选择范围 -- candidates[start: ]
// currentSum: 当前和
// target: 目标和 -- 终止条件
// res: 一次回溯过程中的结果 -- 回溯路径
// result: 总结果集
func dfscombinationSum40(candidates []int, start int, currentSum, target int, res *[]int, result *[][]int) {
	if currentSum == target {
		// 回溯路径已满足条件 -- 将本次回溯的结果计入总结果集后返回
		temp := make([]int, len(*res))
		copy(temp, *res)
		*result = append(*result, temp)
		return
	}

	// 从本层的选择列表中不断固定元素到回溯路径, 然后进行下一层的数字选择并进行回溯
	for i := start; i < len(candidates); i++ {
		if currentSum+candidates[i] > target {
			// 剪枝
			continue
		}

		if i > start && candidates[i] == candidates[i-1] {
			// 剪枝 -- 跳过重复元素
			continue
		}

		// 固定当前层的一个元素到回溯路径
		*res = append(*res, candidates[i])
		// 递归进入下一层选择下一个数字
		// 因为每个数字只能被使用一次，所以进入下一层时当前数字不可以继续使用，因此下一层的start为i+1
		// 如果每个数字可以被使用多次，那么下一层的start应该传i
		dfscombinationSum40(candidates, i+1, currentSum+candidates[i], target, res, result)
		// 将当前元素移出回溯路径 -- 下一次循环选择下一个元素进行固定并继续进入下一层
		*res = (*res)[:len(*res)-1]
	}
}
