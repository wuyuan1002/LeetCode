package main

// 704. 二分查找

// 给定一个n个元素有序的（升序）整型数组nums 和一个目标值target，
// 写一个函数搜索nums中的 target，如果目标值存在返回下标，否则返回 -1。

// func main() {

// }

func search1(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	l, r, mid := 0, len(nums)-1, 0
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		}
	}

	return -1
}
