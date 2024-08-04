package main

// 704. 二分查找

// 给定一个n个元素有序的（升序）整型数组nums 和一个目标值target ，
// 写一个函数搜索nums中的 target，如果目标值存在返回下标，否则返回 -1。

// search .
// leetcode 35. 搜索插入位置
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		// mid := (l + r) / 2
		mid := l + (r-l)/2 // 防止l+r时超出int的范围
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
