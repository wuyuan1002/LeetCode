package main

// 77. 组合

// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//
// 你可以按 任何顺序 返回答案。

// combine .
// leetcode 78. 子集
// leetcode 90. 子集 II
//
// 递归回溯
// 每次选定一个数字后进入下一层继续选择下一个数字到回溯路径，直到回溯路径中的数字个数达到k则说明找到了一个解
// 记录结果后向前回溯，将回溯路径的数字不断重新选择，寻找其他的解并进行记录
//
// 此题也类似动态规划01背包问题 -- leetcode 416. 分割等和子集
// 不过动态规划只能求出解的个数，而无法求出每个解是什么，且此题有重复元素，计算时需要去重
func combine(n int, k int) [][]int {
	if n < k {
		return nil
	}

	res := make([]int, 0)      // 一次回溯过程中的结果 -- 回溯路径
	result := make([][]int, 0) // 总结果集

	dfsCombine(n, 1, k, &res, &result)
	return result
}

// dfsCombine 回溯遍历选择列表，记录满足条件的结果
// n: 选择列表 -- [1, n]
// start: 每次遍历的起始下标 -- 指定当前层的选择范围 -- [start, n]
// k: 终止条件 -- 取k个数的组合
// res: 一次回溯过程中的结果 -- 回溯路径
// result: 总结果集
func dfsCombine(n int, start int, k int, res *[]int, result *[][]int) {
	// 回溯路径已满足条件 -- 将本次回溯的结果计入总结果集后返回
	if len(*res) == k {
		temp := make([]int, k)
		copy(temp, *res)
		*result = append(*result, temp)
		return
	}

	// 遍历选择列表，在当前层选定值并加入到回溯路径中，然后递归进入下一层
	for i := start; i <= n; i++ {
		if len(*res)+n-start+1 < k {
			// 剪枝 -- 若 已有元素+剩余元素 < k 时可以直接跳出，因为无论如何都不会满足总数等于k了
			return
		}

		// 将当前值加入到回溯路径
		*res = append(*res, i)
		// 递归进行下一层
		dfsCombine(n, i+1, k, res, result)
		// 将当前值移出回溯路径
		*res = (*res)[:len(*res)-1]
	}
}
