package main

// 81. 搜索旋转排序数组 II

// 已知存在一个按非降序排列的整数数组 nums ，数组中的值不必互不相同。
//
// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转 ，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,4,4,5,6,6,7] 在下标 5 处经旋转后可能变为 [4,5,6,6,7,0,1,2,4,4] 。
//
// 给你 旋转后 的数组 nums 和一个整数 target ，请你编写一个函数来判断给定的目标值是否存在于数组中。如果 nums 中存在这个目标值 target ，则返回 true ，否则返回 false 。
//
// 你必须尽可能减少整个操作步骤。

// search81 .
// leetcode 33. 搜索旋转排序数组
//
// 二分查找
// 虽然无法通过nums[mid]与target的大小关系确定target在mid的左侧还是右侧，
// 但是可以通过nums[l]、nums[mid]、nums[r]三者来确定好target的方位，然后进行左右指针的收缩，
// 若nums[l]、nums[mid]、nums[r]三者数字相等时，将无法区分mid的左右哪边是已经排好序的，
// 所以此时需要将左右指针分别向中间移动来去除重复数字，在新的区间中进行二分 -- 如[3, 1, 2, 3, 3, 3, 3, 3, 3, 3]
func search81(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target { // 若mid等于目标值则直接返回
			return true
		}

		if nums[l] == nums[mid] && nums[mid] == nums[r] { // 若l、r、mid指向的数字相同，此时无法区分mid的左右哪边是已经排好序的，此时将左右指针分别向中间移动来去除重复数字，在新的区间中进行二分 -- 如[3, 1, 2, 3, 3, 3, 3, 3, 3, 3]
			l++
			r--
		} else if nums[l] <= nums[mid] { // 若mid左面是排好序的 -- 接下来判断应该在mid左面找还是右面找
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

	return false
}
