package main

// 39. 组合总和

// 给你一个 无重复元素 的整数数组candidates 和一个目标整数target，
// 找出candidates中可以使数字和为目标数target 的 所有不同组合 ，
// 并以列表形式返回。你可以按 任意顺序 返回这些组合。
//
// candidates 中的 同一个 数字可以 无限制重复被选取 。
// 如果至少一个数字的被选数量不同，则两种组合是不同的。
//
// 对于给定的输入，保证和为target 的不同组合数少于 150 个。

// combinationSum39 .
// leetcode 40. 组合总和 II
//
// 递归回溯
// 每次选定一个数字后进入下一层继续选择下一个数字到回溯路径，直到总和达到目标值说明找到一个解
// 记录结果后向前回溯，将回溯路径的数字不断重新选择，寻找其他的解并进行记录
//
// 此题也类似动态规划完全背包问题 -- leetcode 518. 零钱兑换
// 不过动态规划只能求出解的个数，而无法求出每个解是什么
func combinationSum39(candidates []int, target int) [][]int {
	if candidates == nil || len(candidates) == 0 {
		return nil
	}

	res := make([]int, 0)      // 一次回溯过程中的结果 -- 回溯路径
	result := make([][]int, 0) // 总结果集

	dfscombinationSum39(candidates, 0, 0, target, &res, &result)
	return result
}

// dfscombinationSum39 回溯遍历选择列表，记录满足条件的结果
// candidates: 选择列表
// start: 每次遍历的起始下标 -- 指定当前层的选择范围
// currentSum: 当前和
// target: 目标和 -- 终止条件
// res: 一次回溯过程中的结果 -- 回溯路径
// result: 总结果集
func dfscombinationSum39(candidates []int, start int, currentSum, target int, res *[]int, result *[][]int) {
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

		// // 此题说明了数组中无重复元素, 所以不用跳过重复元素
		// if i > start && candidates[i] == candidates[i-1] {
		// 	// 剪枝 -- 跳过重复元素
		// 	continue
		// }

		// 将当前值加入到回溯路径
		*res = append(*res, candidates[i])
		// 递归进入下一层选择下一个数字
		// 因为每个数字可以无限制重复被选取，所以进入下一层时当前数字仍然可以被使用，因此下一层的start为i
		// 如果每个数字只能被使用一次，那么下一层的start应该传i+1
		dfscombinationSum39(candidates, i, currentSum+candidates[i], target, res, result)
		// 将当前值移出回溯路径
		*res = (*res)[:len(*res)-1]
	}
}
