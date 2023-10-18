package main

// 47. 全排列 II

// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

// permuteUnique .
// 递归+回溯, 在一层固定一个数字，然后进入下一层固定下一个数字
// 可以使用res回溯列表来存储已经选择好的数字, 也可以使用在原数组中直接交换的方式, 将选好的数字交换到前面来
func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0) // 总结果集

	dfsPermuteUnique(nums, 0, &result)
	return result
}

// dfsPermuteUnique 回溯遍历选择列表，求出nums中数字的全排列
// 可以使用res来存储已经选择好的数字
// 也可以使用直接在原数组中交换数字的方式实现选择数字
//
// nums: 选择列表
// start: 每次遍历的起始下标 -- 指定当前层的选择范围 -- nums[start: ]
// result: 总结果集
func dfsPermuteUnique(nums []int, start int, result *[][]int) {
	// 若nums中所有数字都被选择完了, 说明找到了一种排列方式, 将结果记录到总结果中
	if start == len(nums)-1 {
		if start == len(nums)-1 {
			temp := make([]int, len(nums))
			copy(temp, nums)
			*result = append(*result, temp)
			return
		}
	}

	// 由于本题nums中包含重复数字, 所以需要visited记录已访问的数字, 而46题就不需要
	// 还可以使用先在permuteUnique中给nums排序, 之后在dfs时判断当前数字和前一个数字是否相等来去重
	visited := make(map[int]struct{})

	for i := start; i < len(nums); i++ {
		// 剪枝 -- 去除重复元素, 避免出现重复解 -- 若当前数字在本层已经被选择过则跳过
		if _, ok := visited[nums[i]]; ok {
			continue
		}

		// 标记当前数字在当前层已被选择过
		visited[nums[i]] = struct{}{}

		// 交换当前数字到最前面表示选择当前数字已被选择
		nums[i], nums[start] = nums[start], nums[i]
		// 进入下一层选择下一个数字
		// 当前层选择的数字被交换到start下标处, 下一层开始位置下标为 start+1
		dfsPermuteUnique(nums, start+1, result)
		// 选择当前数字递归结束, 交换回原位置, 继续选择下一个数字进行递归
		nums[i], nums[start] = nums[start], nums[i]
	}
}
