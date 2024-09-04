package main

// 442. 数组中重复的数据

// 给定一个整数数组 a，其中1 ≤ a[i] ≤ n （n为数组长度）, 其中有些元素出现两次而其他元素出现一次。
// 找到所有出现两次的元素。
//
// 你可以不用到任何额外空间并在O(n)时间复杂度内解决这个问题吗？

// findDuplicates .
// leetcode 41. 缺失的第一个正数
// leetcode 287. 寻找重复数
//
// 原地哈希
// 1. 将元素交换到对应的位置（Offer 03）
// 2. 使用正负号作为标记
//
// 使用 下标k 处数字的正负号表示 数字k+1 是否在数组中出现过
// 一次遍历数组，若下标处数字为正，说明该数字是第一次出现，将下标处数字置为负数，
// 若下标处数字为负数，说明当前数字已经出现过，为重复数字，将数字写入结果集
func findDuplicates(nums []int) []int {
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

// abs .
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
