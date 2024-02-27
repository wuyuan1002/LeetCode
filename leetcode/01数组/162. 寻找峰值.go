package main

// 162. 寻找峰值

// 峰值元素是指其值严格大于左右相邻值的元素。
//
// 给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
//
// 你可以假设 nums[-1] = nums[n] = -∞ 。
//
// 你必须实现时间复杂度为 O(log n) 的算法来解决此问题。

// findPeakElement .
// 1. 一次遍历数组，寻找比它前一个和后一个都大的元素下标 O(n)
// 2. 二分查找 O(log n) -- 只要数组中存在一个元素比相邻元素大，那么沿着它一定可以找到一个峰值
func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2 // 防止l+r时超出int的范围
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
