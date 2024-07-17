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
// leetcode 852. 山脉数组的峰顶索引
// leetcode 153. 寻找旋转排序数组中的最小值
//
// 二分查找 -- 只要数组中存在一个元素比相邻元素大，那么沿着它一定可以找到一个峰值
// 这里需要使用 l < r 的二分查找，而不能使用 l <= r 的二分查找，
// 当nums[mid]比nums[mid+1]大时，说明mid到r是单调递减的，最大值一定在mid左侧（因为可能正好是mid，所以这里不能用 l <= r 的那种循环，也就不能使用 r=mid-1 这样赋值）
// 否则说明0到mid+1是单调递增的，最大值一定出现在mid右侧（因为0到mid是单调递增的且求的是最大值，此时mid一定不会是最大值，可能是mid+1，所以可以直接 l=mid+1 赋值）
func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2

		// 若nums[mid] > nums[mid+1]，说明峰值位于nums[0, mid]范围内（可能正好是nums[mid]），所以将r指向mid
		// 若nums[mid] <= nums[r]，说明峰值位于nums(mid, r]范围内（一定不会是nums[mid]），所以将l指向mid+1
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	// 二分结束后，左指针指向的就是一个峰值
	return l
}
