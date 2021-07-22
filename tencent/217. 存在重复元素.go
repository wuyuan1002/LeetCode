package main

// 217. 存在重复元素

// 给定一个整数数组，判断是否存在重复元素。
//
// 如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。

func main() {

}

// 1. 排序
// 2. map
func containsDuplicate(nums []int) bool {
	if nums == nil || len(nums) < 2 {
		return false
	}

	hash := make(map[int]bool)
	for _, num := range nums {
		if hash[num] {
			return true
		}
		hash[num] = true
	}

	return false
}
