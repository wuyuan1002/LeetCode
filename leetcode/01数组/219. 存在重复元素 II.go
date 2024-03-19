package main

// 219. 存在重复元素 II

// 给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，
// 满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。

// containsNearbyDuplicate .
// 滑动窗口 -- 维护长度为k的窗口，每次向右移动都判断添加的元素在窗口中是否有重复
func containsNearbyDuplicate(nums []int, k int) bool {

	// 使用一个map作为滑动窗口，因为窗口向右移动时需要判断添加元素是否在窗口中存在 -- map中最多存储k个数字
	hash := make(map[int]bool, k)

	for i, n := range nums {
		// 若窗口中元素大于k了，则先将窗口左侧元素移出
		if i > k {
			delete(hash, nums[i-k-1])
		}

		// 若新加入窗口的数字在窗口中已出现，则返回true
		if _, ok := hash[n]; ok {
			return true
		}

		// 新的数字在窗口中首次出现 -- 加入到窗口中
		hash[n] = true
	}

	// 若遍历完数组都没有在窗口中发现重复数字则返回false
	return false
}
