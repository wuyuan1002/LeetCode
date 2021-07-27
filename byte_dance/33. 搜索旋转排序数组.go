package main

import (
	"fmt"
)

// 33. 搜索旋转排序数组

// 整数数组 nums 按升序排列，数组中的值互不相同 。
// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了旋转，
// 使数组变为 [nums[k], nums[k+1], ..., nums[n-1],
// nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
// 例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为[4,5,6,7,0,1,2] 。
//
// 给你旋转后的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，
// 则返回它的下标，否则返回-1。

func main() {
	fmt.Println(search([]int{1}, 3))
}

// 二分法
func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[0] <= nums[mid] { // 若左面是排好序的
			if nums[0] <= target && nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[len(nums)-1] >= target && nums[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}

// 二分法
// 第二次做
func search2(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[0] <= nums[mid] { // 左面是排好序的
			if nums[0] <= target && nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[len(nums)-1] >= target && nums[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}
