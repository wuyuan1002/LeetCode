package main

// 33. 搜索旋转排序数组

// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了旋转，
// 使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标从0开始计数）。
// 例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为[4,5,6,7,0,1,2] 。
//
// 给你旋转后的数组nums和一个整数target，如果 nums 中存在这个目标值target，则返回它的下标，否则返回-1。

// search33 .
// leetcode 81. 搜索旋转排序数组 II
//
// 二分查找
// 虽然无法通过nums[mid]与target的大小关系确定target在mid的左侧还是右侧，
// 但是可以通过nums[l]、nums[mid]、nums[r]三者来确定好target的方位，然后进行左右指针的收缩
func search33(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target { // 若mid等于目标值则直接返回
			return mid
		}

		if nums[l] <= nums[mid] { // 若mid左面是排好序的 -- 接下来判断应该在mid左面找还是右面找
			if nums[mid] > target && nums[l] <= target {
				// 若目标值在mid左面排好序的区间内 [l, mid) ，说明应该在mid左面寻找
				r = mid - 1
			} else {
				// 否则应该在mid右面寻找
				l = mid + 1
			}
		} else { // 若mid右面是排好序的 -- 接下来判断应该在mid左面找还是右面找
			if nums[mid] < target && nums[r] >= target {
				// 若目标值在mid右面排好序的区间内 (mid, r] ，说明应该在mid右面寻找
				l = mid + 1
			} else {
				// 否则应该在mid左面寻找
				r = mid - 1
			}
		}
	}

	return -1
}
