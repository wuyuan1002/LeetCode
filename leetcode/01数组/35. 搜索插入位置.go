package main

// 35. 搜索插入位置

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
// 如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
// 请必须使用时间复杂度为 O(log n) 的算法。

// searchInsert 二分查找target，若找到则直接返回，若没找到则左指针的位置就是它的插入位置
// offer 11、leetcode 704、33、35、153
func searchInsert(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return mid
		}
	}
	return l
}
