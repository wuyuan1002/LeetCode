package main

// 34. 在排序数组中查找元素的第一个和最后一个位置

// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
// 如果数组中不存在目标值 target，返回[-1, -1]。
// 进阶：
// 你可以设计并实现时间复杂度为O(log n)的算法解决此问题吗？

// searchRange .
// 二分查找
// 分别二分查找目标值在数组中的第一个值和最后一个值，然后返回两个坐标
func searchRange(nums []int, target int) []int {
	return []int{getFirst(nums, target), getLast(nums, target)}
}

// getFirst 获取第一个
func getFirst(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			// 若找到target但是其左侧仍然存在与目标值相同的值，说明当前找到的不是第一个target，需要继续在左侧寻找
			if mid > 0 && nums[mid-1] == target {
				r = mid - 1
			} else {
				return mid
			}
		}
	}
	return -1
}

// getLast 获取最后一个
func getLast(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			// 若找到target但是其右侧仍然存在与目标值相同的值，说明当前找到的不是最后一个target，需要继续在右侧寻找
			if mid < len(nums)-1 && nums[mid+1] == target {
				l = mid + 1
			} else {
				return mid
			}
		}
	}
	return -1
}
