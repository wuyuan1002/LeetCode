package main

// 40. 组合总和 II

// 给定一个数组candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。
// candidates中的每个数字在每个组合中只能使用一次。
//
// 注意：解集不能包含重复的组合。

// combinationSum2
// 1. 回溯法
// 2. 动态规划 -- 01背包问题 -- 此题有重复元素，计算时无法去重
func combinationSum2(candidates []int, target int) [][]int {
	if candidates == nil || len(candidates) == 0 {
		return nil
	}

	// 先排序方便剪枝
	// 也可在dfs时使用visited记录已访问的数字来进行剪枝去重
	quickSort40(candidates, 0, len(candidates)-1)

	res := make([]int, 0)      // 一次回溯过程中的结果 -- 回溯路径
	result := make([][]int, 0) // 总结果集
	dfsCombinationSum2(candidates, 0, 0, target, &res, &result)
	return result
}

// dfsCombinationSum2 回溯遍历选择列表，记录满足条件的结果
// // candidates: 选择列表
// // start: 每次遍历的起始下标 -- 指定当前层的选择范围
// // currentSum: 当前和
// // target: 目标和 -- 终止条件
// // res: 一次回溯过程中的结果 -- 回溯路径
// // result: 总结果集
func dfsCombinationSum2(candidates []int, start int, currentSum, target int, res *[]int, result *[][]int) {
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
		// 递归进入下一层 -- 因为每个数字只能使用一次,
		// 所以进入下一层时当前数字不可以继续使用, 因此下一层的start为 i+1
		dfsCombinationSum2(candidates, i+1, currentSum+candidates[i], target, res, result)
		// 将当前元素移出回溯路径 -- 下一次循环选择下一个元素进行固定并继续进入下一层
		*res = (*res)[:len(*res)-1]
	}
}

// quickSort40 快排
func quickSort40(nums []int, left, right int) {
	if nums == nil || len(nums) == 0 || left >= right {
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

	nums[l], nums[left] = nums[left], nums[l]
	quickSort40(nums, left, l-1)
	quickSort40(nums, r+1, right)
}
