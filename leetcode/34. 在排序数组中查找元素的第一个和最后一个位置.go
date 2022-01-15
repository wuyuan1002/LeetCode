package main

// 34. 在排序数组中查找元素的第一个和最后一个位置

// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
// 如果数组中不存在目标值 target，返回[-1, -1]。
//
// 进阶：
// 你可以设计并实现时间复杂度为O(log n)的算法解决此问题吗？

// func main() {

// }

func searchRange(nums []int, target int) []int {
	return []int{getFirst(nums, target), getLast(nums, target)}
}

// 获取第一个
func getFirst(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			if mid > 0 && nums[mid-1] == target {
				r = mid - 1
			} else {
				return mid
			}
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

// 获取最后一个
func getLast(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			if mid < len(nums)-1 && nums[mid+1] == target {
				l = mid + 1
			} else {
				return mid
			}
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
