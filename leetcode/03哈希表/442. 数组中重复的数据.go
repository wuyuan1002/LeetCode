package main

// 442. 数组中重复的数据

// 给定一个整数数组 a，其中1 ≤ a[i] ≤ n （n为数组长度）, 其中有些元素出现两次而其他元素出现一次。
// 找到所有出现两次的元素。
//
// 你可以不用到任何额外空间并在O(n)时间复杂度内解决这个问题吗？

// findDuplicates1 .
// leetcode 41. 缺失的第一个正数
// leetcode 448. 找到所有数组中消失的数字
// leetcode 287. 寻找重复数
//
// 原地哈希
// 1. 使用正负号作为标记 -- 使用 下标k 处数字的正负号表示 数字k+1 是否在数组中出现过
//
// 一次遍历数组，若下标处数字为正，说明该数字是第一次出现，将下标处数字置为负数，
// 若下标处数字为负数，说明当前数字已经出现过，为重复数字，将数字写入结果集
func findDuplicates1(nums []int) []int {
	result := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		// 获取当前数字的绝对值 -- 因为当前位置数字可能为负数, 所以要使用绝对值
		num := abs(nums[i])

		// 若当前数字对应的下标处数字为负数，说明当前数字已经出现过，将其记入结果集，
		// 否则说明当前数字首次出现，将下标处数字置为负数
		if nums[num-1] < 0 {
			result = append(result, num)
		} else {
			nums[num-1] = -nums[num-1]
		}
	}

	return result
}

// findDuplicates2 .
// Offer 03. 数组中重复的数字
//
// 2. 将元素交换到对应的位置 -- 将 数字n 不断归位到 下标n-1 处，已在正确位置的数字表示数字出现过
//
// 一次遍历数组，不断将 数字n 归位到 下标n-1 处，若发现要交换的数字已经在正确的位置，
// 说明这个数字重复了，将其记入结果集（记录总结果时需要进行去重）
func findDuplicates2(nums []int) []int {
	// 临时存放结果，用于去重 -- 因为重复的数字会被遍历到多次，所以需要去重
	unique := make(map[int]struct{})

	// 不断遍历数组的每一个数字将其归位到正确的位置上，若发现重复则将其记入结果集
	for i, n := 0, 0; i < len(nums); {
		// 获取当前下标的数字
		n = nums[i]

		if n-1 != i {
			// 若当前数字不在正确的位置上，则将它交换到与其相等的下标处

			if n != nums[n-1] {
				// 若若目标交换下标处的数字不在正确位置，则将当前数字与其交换归位到正确位置，并继续检查当前下标处新交换过来的数字是否在正确位置
				nums[i], nums[n-1] = nums[n-1], nums[i]
			} else {
				// 若目标交换下标处的数字已经在正确位置，说明当前数字是重复数字，将其记入结果集并继续遍历下一个下标
				unique[n] = struct{}{}
				i++
			}
		} else {
			// 若当前下标处数字已在正确位置，则继续遍历数组下一个位置
			i++
		}
	}

	// 遍历结果，构造返回数据
	result := make([]int, 0, len(unique))
	for n := range unique {
		result = append(result, n)
	}

	return result
}

// abs .
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
