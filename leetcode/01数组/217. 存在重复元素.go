package main

// 217. 存在重复元素

// 给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；
// 如果数组中每个元素互不相同，返回 false 。

// containsDuplicate .
// 1. 排序
// 2. map
func containsDuplicate(nums []int) bool {
	hash := make(map[int]bool)
	for _, num := range nums {
		if hash[num] {
			return true
		}
		hash[num] = true
	}

	return false
}
