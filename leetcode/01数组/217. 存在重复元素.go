package main

// 217. 存在重复元素

// 给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；
// 如果数组中每个元素互不相同，返回 false 。

// containsDuplicate .
// 从头遍历数组元素，将访问到的数字存到map中，若在存入map时发现已经存在则说明该数字出现重复
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
