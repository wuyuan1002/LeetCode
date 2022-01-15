package main

// 162. 寻找峰值

// 峰值元素是指其值大于左右相邻值的元素。
//
// 给你一个输入数组nums，找到峰值元素并返回其索引。数组可能包含多个峰值，
// 在这种情况下，返回 任何一个峰值 所在位置即可。
//
// 你可以假设nums[-1] = nums[n] = -∞ 。

// func main() {

// }

// 1. 一次遍历数组，寻找比它前一个和后一个都大的元素下标 O(n)
// 2. 二分查找 O(log n)
func findPeakElement(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
