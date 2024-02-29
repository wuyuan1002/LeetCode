package main

// 216. 组合总和 III

// 找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
//
// 只使用数字1到9
// 每个数字 最多使用一次
// 返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。

// combinationSum3 .
func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)
	res := make([]int, 0)
	dfsCombinationSum3(1, 0, n, k, &res, &result)
	return result
}

// dfsCombinationSum3 回溯遍历选择列表，记录满足条件的结果
// start: 选择列表 -- 即本层可以选择的数字在 1-9 中的起始下标
// currentSum: 当前和
// target: 目标和 -- 终止条件
// k: 组成一个结果的元素个数限制 -- 终止条件
// res: 一次回溯过程中的结果 -- 回溯路径
// result: 总结果集
func dfsCombinationSum3(start int, currentSum, target, k int, res *[]int, result *[][]int) {
	// 若当前和等于目标和且组成目标和的元素个数也满足条件则记录当前回溯路径的结果到总结果集
	if currentSum == target && len(*res) == k {
		tmp := make([]int, len(*res))
		copy(tmp, *res)
		*result = append(*result, tmp)
		return
	}

	// 从本层的选择列表中不断固定元素到回溯路径, 然后进行下一层的数字选择并进行回溯
	for i := start; i <= 9; i++ {
		sum := currentSum + i
		if sum > target || len(*res) >= k {
			// 剪枝
			return
		}

		// 固定当前元素到选择列表并进入下一层进行数字选择和回溯
		*res = append(*res, i)
		dfsCombinationSum3(i+1, sum, target, k, res, result)
		*res = (*res)[:len(*res)-1]
	}
}
