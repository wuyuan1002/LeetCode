package main

// 35. 搜索插入位置

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
// 如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
// 请必须使用时间复杂度为 O(log n) 的算法。

// searchInsert .
// leetcode 704、33、35、153、offer 11
//
// 二分查找 -- 若找到目标值则直接返回，若没找到则左指针的位置就是它的插入位置（因为返回时左指针已经向右移动过一位了，此位置就是插入位置）
//
// 1. 若找到目标值 -- 则直接返回其下标
// 2. 若未找到目标值 -- 则返回左指针下标
// 因为如果上面的没有返回 return mid，说明最后一定是由于 left > right 从而跳出循环的，在此之前是 left == right（最后一轮循环时），
// - 如果最后是 right-1 导致的 left > right，说明原来的right位置是大于target的，所以返回原来的right位置即left的位置
// - 如果最后是 left+1 导致的 left > right，说明是原来的left位置是小于target的，而right能移动到这个位置，说明此位置右侧数字都是大于target的，left + 1 就正好移动到了第一个大于target的数字位置，返回加1后的left即可
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
