package main

import "sort"

// 658. 找到 K 个最接近的元素

// 给定一个 排序好 的数组 arr ，两个整数 k 和 x ，从数组中找到最靠近 x（两数之差最小）的 k 个数。返回的结果必须要是按升序排好的。
//
// 整数 a 比整数 b 更接近 x 需要满足：
// |a - x| < |b - x| 或者
// |a - x| == |b - x| 且 a < b

// findClosestElements .
// 若差值不同，则差值越小越靠近，若差值相同则数字越小越靠近
// 1. 将数组按照「更靠近」进行重新排序，排序后数组的前k个数字就是答案 -- 由于数组本身就是从小到大排序好的，所以若差值相同时，本身在前的更靠近，因此使用稳定排序即可在差值相同时不破坏原有顺序
func findClosestElements(arr []int, k int, x int) []int {
	// 稳定排序，在差值绝对值相同的情况下，保证更小的数排在前面
	sort.SliceStable(arr, func(i, j int) bool { return abs(arr[i]-x) < abs(arr[j]-x) })
	arr = arr[:k]
	sort.Ints(arr)
	return arr
}

// findClosestElements .
// 若差值不同，则差值越小越靠近，若差值相同则数字越小越靠近
// 2. 先使用二分查找找到数组中最接近x的数字，然后使用双指针向两边进行拓展
func findClosestElements1(arr []int, k int, x int) []int {
	right := sort.SearchInts(arr, x) // 找到第一个大于等于x的数字下标
	left := right - 1                // 最后一个小于x的数字下标

	// 双指针向两边进行拓展直到区间大小扩展到k
	for ; k > 0; k-- {
		if left < 0 {
			right++
		} else if right >= len(arr) || x-arr[left] <= arr[right]-x {
			// 若左右都可扩展，且左右差值相同时，优先想左侧扩展
			left--
		} else {
			right++
		}
	}

	return arr[left+1 : right]
}
